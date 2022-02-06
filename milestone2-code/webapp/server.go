package main

import (
	"fmt"
	"net/http"

	"github.com/ivivanov18/WebAppGoScaffold/milestone2-code/webapp/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "You have reached Echorand Corp’s Service API")
	})
	mux.HandleFunc("/static", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "static file")
	})
	mux.HandleFunc("/healthcheck", handlers.HealthCheck)
	mux.HandleFunc("/", handlers.Index)

	http.ListenAndServe(":8080", mux)
}
