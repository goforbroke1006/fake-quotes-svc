package svc_http

import (
	"go.uber.org/atomic"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func New(isReady *atomic.Bool) *http.ServeMux {
	mux := http.ServeMux{}
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("healthy"))
	})
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, req *http.Request) {
		if isReady.Load() {
			w.WriteHeader(200)
			_, _ = w.Write([]byte("ready"))
		} else {
			w.WriteHeader(404)
			_, _ = w.Write([]byte("is not ready"))
		}
	})
	mux.Handle("/metrics", promhttp.Handler())
	return &mux
}
