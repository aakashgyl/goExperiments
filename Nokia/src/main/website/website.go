package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/bar/", handleReq)

	http.ListenAndServe(":8080", nil)
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL)
}
