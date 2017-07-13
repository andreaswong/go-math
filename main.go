package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/api/math/even", func(w http.ResponseWriter, r *http.Request) {
		numberStr := r.URL.Query().Get("number")

		if len(numberStr) > 0 {
			if number, err := strconv.Atoi(numberStr); err == nil {
				if number%2 == 0 {
					fmt.Fprintf(w, "true")
				} else {
					fmt.Fprintf(w, "false")
				}
			} else {
				fmt.Fprintf(w, numberStr + " is not a number")
			}
		}
	})

	fmt.Println(http.ListenAndServe(":8889", nil))
}
