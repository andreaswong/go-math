package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/api/math/even", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":8280", nil)
}