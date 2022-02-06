package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "Static page Index")
}
