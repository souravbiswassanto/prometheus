package api

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping Handler",
	})

func ping(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
}
func RunServer(Port string) {
	prometheus.MustRegister(pingCounter)
	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+Port, nil)
}
