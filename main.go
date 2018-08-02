package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8889"
	}

	http.Handle("/", R)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}
