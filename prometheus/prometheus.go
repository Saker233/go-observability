package prometheus

import (
	"net/http"

	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Registry = prom.NewRegistry()

	WorkRequests = prom.NewCounterVec(
		prom.CounterOpts{
			Name: "work_requests_total",
			Help: "Total requests to /work",
		},
		[]string{"status"},
	)

	WorkDuration = prom.NewHistogram(
		prom.HistogramOpts{
			Name:    "work_duration_seconds",
			Help:    "Duration of /work requests in seconds",
			Buckets: prom.DefBuckets,
		},
	)

	WorkInFlight = prom.NewGauge(
		prom.GaugeOpts{
			Name: "work_requests_in_flight",
			Help: "Currently executing /work requests",
		},
	)
)

func init() {
	Registry.MustRegister(
		WorkRequests,
		WorkDuration,
		WorkInFlight,
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(
			collectors.ProcessCollectorOpts{},
		),
	)
}

func Handler() http.Handler {
	return promhttp.HandlerFor(
		Registry,
		promhttp.HandlerOpts{},
	)
}
