package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func GetRouter() *mux.Router {
	R := mux.NewRouter()
	R.Methods("GET").Path("/api/math/even").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqDTO := mux.Vars(r)

		var numberStr string
		values, ok := r.URL.Query()["number"]
		if ok {
			numberStr = values[0]
		}

		logrus.Infof("req=%#v", reqDTO)
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

	return R
}
