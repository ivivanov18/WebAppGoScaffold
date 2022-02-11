package main

import (
	"fmt"
	"net/http"

	"github.com/ivivanov18/WebAppGoScaffold/milestone2-code/webapp/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "You have reached Echorand Corp’s Service API")
	})
	mux.HandleFunc("/healthcheck", handlers.HealthCheck)
	http.ListenAndServe(":8080", mux)
}
