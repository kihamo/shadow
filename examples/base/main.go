package main // import "github.com/kihamo/shadow/examples/base"

import (
	"log"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/resource/alerts"
	"github.com/kihamo/shadow/resource/config"
	"github.com/kihamo/shadow/resource/logger"
	"github.com/kihamo/shadow/resource/mail"
	"github.com/kihamo/shadow/resource/metrics"
	"github.com/kihamo/shadow/resource/template"
	"github.com/kihamo/shadow/resource/workers"
	"github.com/kihamo/shadow/service/frontend"
	"github.com/kihamo/shadow/service/system"
)

func main() {
	application, err := shadow.NewApplication(
		[]shadow.Resource{
			new(config.Config),
			new(logger.Logger),
			new(template.Template),
			new(workers.Workers),
			new(mail.Mail),
			new(metrics.Metrics),
			new(alerts.Alerts),
		},
		[]shadow.Service{
			new(system.SystemService),
			new(frontend.FrontendService),
		},
		"Shadow base",
		"1.0",
		"12345-base",
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = application.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
