package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const pushgatewayAddr = "http://pushgateway:9091"
const applicationPort = ":2112"

func main() {
	println("starting server...")
	version := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "version",
		Help: "Version information about this service",
		ConstLabels: map[string]string{
			"version": Version,
			"service": ApplicationName,
		},
	})
	prometheus.MustRegister(version)
	// Easy case:
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "counter_app",
		Help: "Counter test",
	})
	prometheus.MustRegister(counter)
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		counter.Inc()
		handler := promhttp.Handler()
		handler.ServeHTTP(w, r)
	})
	if err := http.ListenAndServe(applicationPort, nil); err != nil {
		panic(err)
	}
}
