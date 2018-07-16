package metrics

import (
	"github.com/kihamo/snitch"
)

type HasMetrics interface {
	Metrics() snitch.Collector
}

var (
	// For all responses
	MetricResponseTimeSeconds        = snitch.NewTimer(MetricNameResponseTimeSeconds, "Response time in seconds")
	MetricResponseSizeBytes          = snitch.NewHistogram(MetricNameResponseSizeBytes, "Size of responses in bytes")
	MetricResponseMarshalTimeSeconds = snitch.NewHistogram(MetricNameResponseMarshalTimeSeconds, "Marshal time of response in seconds")

	// For all requests
	MetricRequestSizeBytes                = snitch.NewHistogram(MetricNameRequestSizeBytes, "Size of requests in bytes")
	MetricRequestsTotal                   = snitch.NewCounter(MetricNameRequestsTotal, "Requests total")
	MetricRequestReadTimeSeconds          = snitch.NewTimer(MetricNameRequestReadTimeSeconds, "Read time of request in seconds")
	MetricRequestReadUnmarshalTimeSeconds = snitch.NewTimer(MetricNameRequestReadUnmarshalTimeSeconds, "Read unmarshal time of request in seconds")
	MetricRequestUnmarshalTimeSeconds     = snitch.NewTimer(MetricNameRequestUnmarshalTimeSeconds, "Unmarshal time of request in seconds")

	// For all external requests
	MetricExternalResponseTimeSeconds = snitch.NewTimer(MetricNameExternalResponseTimeSeconds, "External response time in total")
)
