package internal

import (
	"strings"
	"sync"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/i18n"
	"github.com/kihamo/shadow/components/logging"
	"github.com/kihamo/shadow/components/mail"
	"github.com/kihamo/shadow/components/metrics"
	"gopkg.in/gomail.v2"
)

const (
	mailDaemonTimeOut = 5 * time.Minute
)

type mailTask struct {
	message *gomail.Message
	result  chan error
}

type Component struct {
	config config.Component
	logger logging.Logger
	routes []dashboard.Route

	mutex  sync.RWMutex
	open   bool
	dialer *gomail.Dialer
	closer gomail.SendCloser
	queue  chan *mailTask
}

func (c *Component) Name() string {
	return mail.ComponentName
}

func (c *Component) Version() string {
	return mail.ComponentVersion
}

func (c *Component) Dependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name:     config.ComponentName,
			Required: true,
		},
		{
			Name: dashboard.ComponentName,
		},
		{
			Name: i18n.ComponentName,
		},
		{
			Name: logging.ComponentName,
		},
		{
			Name: metrics.ComponentName,
		},
	}
}

func (c *Component) Run(a shadow.Application, ready chan<- struct{}) error {
	c.open = false
	c.queue = make(chan *mailTask)
	c.logger = logging.DefaultLogger().Named(c.Name())
	metricsEnabled := a.HasComponent(metrics.ComponentName)

	<-a.ReadyComponent(config.ComponentName)
	c.config = a.GetComponent(config.ComponentName).(config.Component)

	c.initDialer(
		c.config.String(mail.ConfigSmtpHost),
		c.config.Int(mail.ConfigSmtpPort),
		c.config.String(mail.ConfigSmtpUsername),
		c.config.String(mail.ConfigSmtpPassword),
	)

	ready <- struct{}{}

	for {
		select {
		case task, ok := <-c.queue:
			if !ok {
				return nil
			}

			err := c.execute(task)
			if metricsEnabled {
				metricMailTotal.Inc()

				if err != nil {
					metricMailTotal.With("status", "failed").Inc()
				} else {
					metricMailTotal.With("status", "success").Inc()
				}
			}

		case <-time.After(mailDaemonTimeOut):
			c.mutex.Lock()
			if c.open {
				if err := c.closer.Close(); err != nil && !strings.Contains(err.Error(), "4.4.2") {
					c.logger.Error("Dialer close failed", "error", err.Error())
				} else {
					c.logger.Debug("Dialer close success")
				}

				c.open = false
			}
			c.mutex.Unlock()
		}
	}

	return nil
}

func (c *Component) initDialer(host string, port int, username, password string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.dialer = gomail.NewDialer(host, port, username, password)
	c.open = false
}

func (c *Component) execute(task *mailTask) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var err error

	if !c.open {
		if c.closer, err = c.dialer.Dial(); err != nil {
			c.logger.Error("Dialer dial failed", "error", err.Error())
			task.result <- err

			return err
		} else {
			c.logger.Debug("Dialer open success")
			c.open = true
		}
	}

	if c.open {
		if len(task.message.GetHeader("From")) == 0 {
			task.message.SetAddressHeader("From", c.config.String(mail.ConfigFromAddress), c.config.String(mail.ConfigFromName))
		}

		if err = gomail.Send(c.closer, task.message); err != nil {
			if strings.Contains(err.Error(), "4.4.2") {
				c.logger.Debug("SMTP server response timeout exceeded",
					"mail", task.message,
					"error", err.Error(),
				)

				c.open = false
				return c.execute(task)
			} else {
				c.logger.Error(err.Error(), "mail", task.message)
				task.result <- err

				return err
			}
		} else {
			c.logger.Debug("Send message success", "mail", task.message)
			task.result <- nil
		}
	}

	return nil
}

func (c *Component) Send(message *gomail.Message) {
	task := &mailTask{
		message: message,
		result:  make(chan error),
	}
	c.queue <- task

	c.logger.Debug("Send new message to queue", "mail", message)

}

func (c *Component) SendAndReturn(message *gomail.Message) error {
	task := &mailTask{
		message: message,
		result:  make(chan error),
	}
	c.queue <- task

	return <-task.result
}
