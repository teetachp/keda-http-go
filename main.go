package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requests",
		Help: "number of requests",
	})
)

func main() {
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		defer httpRequestsCounter.Inc()
		w.Write([]byte("Demo...."))
	})

	http.ListenAndServe(":8080", nil)
}
