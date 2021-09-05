package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createCounter prometheus.Counter
	updateCounter prometheus.Counter
	removeCounter prometheus.Counter
)

func RegisterMetrics() {
	createCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_timeline_api_create",
	})
	updateCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_timeline_api_update",
	})
	removeCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_timeline_api_remove",
	})
}

func IncCreateCounter() {
	createCounter.Inc()
}

func IncUpdateCounter() {
	updateCounter.Inc()
}

func IncRemoveCounter() {
	removeCounter.Inc()
}
