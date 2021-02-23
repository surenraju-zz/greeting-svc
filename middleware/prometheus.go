package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "greeting_svc_http_requests_total",
			Help: "How many HTTP requests processed, partitioned by path.",
		},
		[]string{"path"},
	)
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "greeting_svc_http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
)

func init() {
	prometheus.MustRegister(httpReqs)
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		httpReqs.WithLabelValues(path).Inc()
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}
