package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	p, err := LoadPage("index.html")
	if err != nil {
		fmt.Fprintf(w, "invalid file")
	}
	fmt.Fprint(w, string(p.Body))
}
