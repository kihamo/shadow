package mail

import (
	"strings"
	"sync"
	"time"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/logger"
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
	application shadow.Application
	config      *config.Component
	logger      logger.Logger

	mutex  sync.RWMutex
	open   bool
	dialer *gomail.Dialer
	closer gomail.SendCloser
	queue  chan *mailTask
}

func (c *Component) GetName() string {
	return "mail"
}

func (c *Component) GetVersion() string {
	return "1.0.0"
}

func (c *Component) Init(a shadow.Application) error {
	resourceConfig, err := a.GetComponent("config")
	if err != nil {
		return err
	}
	c.config = resourceConfig.(*config.Component)

	c.application = a

	return nil
}

func (c *Component) Run(wg *sync.WaitGroup) error {
	c.logger = logger.NewOrNop(c.GetName(), c.application)

	c.open = false
	c.queue = make(chan *mailTask)
	c.initDialer(
		c.config.GetString(ConfigMailSmtpHost),
		c.config.GetInt(ConfigMailSmtpPort),
		c.config.GetString(ConfigMailSmtpUsername),
		c.config.GetString(ConfigMailSmtpPassword),
	)

	go func() {
		defer wg.Done()

		for {
			select {
			case task, ok := <-c.queue:
				if !ok {
					return
				}

				err := c.execute(task)
				if metricMailTotal != nil {
					if err != nil {
						metricMailTotal.With("result", "failed").Add(1)
					} else {
						metricMailTotal.With("result", "success").Add(1)
					}
				}

			case <-time.After(mailDaemonTimeOut):
				if c.open {
					if err := c.closer.Close(); err != nil && !strings.Contains(err.Error(), "4.4.2") {
						c.logger.Error("Dialer close failed", map[string]interface{}{"error": err.Error()})
					} else {
						c.logger.Debug("Dialer close success")
					}

					c.open = false
				}
			}
		}
	}()

	return nil
}

func (c *Component) initDialer(host string, port int, username, password string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.dialer = gomail.NewDialer(host, port, username, password)
	c.open = false
}

func (c *Component) execute(task *mailTask) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	var err error

	if !c.open {
		if c.closer, err = c.dialer.Dial(); err != nil {
			c.logger.Error("Dialer dial failed", map[string]interface{}{"error": err.Error()})
			task.result <- err

			return err
		} else {
			c.logger.Debug("Dialer open success")
			c.open = true
		}
	}

	if c.open {
		if len(task.message.GetHeader("From")) == 0 {
			task.message.SetAddressHeader("From", c.config.GetString(ConfigMailFromAddress), c.config.GetString(ConfigMailFromName))
		}

		if err = gomail.Send(c.closer, task.message); err != nil {
			if strings.Contains(err.Error(), "4.4.2") {
				c.logger.Debug("SMTP server response timeout exceeded", map[string]interface{}{
					"message": task.message,
					"error":   err.Error(),
				})

				c.open = false
				return c.execute(task)
			} else {
				c.logger.Error(err.Error(), map[string]interface{}{"message": task.message})
				task.result <- err

				return err
			}
		} else {
			c.logger.Debug("Send message success", map[string]interface{}{"message": task.message})
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

	c.logger.Debug("Send new message to queue", map[string]interface{}{"message": message})

}

func (c *Component) SendAndReturn(message *gomail.Message) error {
	task := &mailTask{
		message: message,
		result:  make(chan error),
	}
	c.queue <- task

	return <-task.result
}