package metrics

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var request = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Namespace:  "metrics",
	Subsystem:  "http",
	Name:       "request",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
}, []string{"status"})

func ObserveRequest(d time.Duration, status int) {
	request.WithLabelValues(strconv.Itoa(status)).Observe(d.Seconds())
}
