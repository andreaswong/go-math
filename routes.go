package main

import "github.com/gorilla/mux"

func init() {
	r := mux.NewRouter()
	r.Get("/api/math/even").HandlerF
}
