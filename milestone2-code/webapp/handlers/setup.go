package handlers

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/", Index)
	mux.Handle("/static/", fs)
	mux.HandleFunc("/api", Api)
	mux.HandleFunc("/healthcheck", HealthCheck)
}
