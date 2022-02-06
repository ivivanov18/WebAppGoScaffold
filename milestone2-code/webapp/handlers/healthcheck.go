package handlers

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok")
}
