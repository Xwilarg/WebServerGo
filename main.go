package main

import (
	"fmt"
	"log"
	"net/http"
)

func getRequest(rw http.ResponseWriter, http *http.Request) {
	switch http.Method {
	case "GET":
		fmt.Fprintf(rw, "OK")
	case "POST":
		fmt.Fprintf(rw, "POST")
	}
}

func main() {
    http.HandleFunc("/", getRequest)
    log.Fatal(http.ListenAndServe(":8082", nil))
}