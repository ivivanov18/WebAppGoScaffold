package main

import (
	"net/http"

	"github.com/ivivanov18/WebAppGoScaffold/milestone2-code/webapp/handlers"
)

func main() {
	mux := http.NewServeMux()
	handlers.RegisterRoutes(mux)
	http.ListenAndServe(":8080", mux)
}
