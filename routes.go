package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var R *mux.Router

func init() {
	R = mux.NewRouter()
	R.Get("/api/math/even").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqDTO := mux.Vars(r)
		logrus.Info("req=%#v", reqDTO)
		numberStr := reqDTO["number"]

		if len(numberStr) > 0 {
			if number, err := strconv.Atoi(numberStr); err == nil {
				if number%2 == 0 {
					fmt.Fprintf(w, "true")
				} else {
					fmt.Fprintf(w, "false")
				}
			} else {
				fmt.Fprintf(w, numberStr+" is not a number")
			}
		}
	}))
}
