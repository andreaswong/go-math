package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "gogogo", nil)
	})

	http.HandleFunc("/api/math/even", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":80", nil)
}