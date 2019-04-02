package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"bytes"
	"strconv"
)

var colors []string

func setHeaders(rw http.ResponseWriter) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Content-Type", "application/json")
}

func getColors(rw http.ResponseWriter, req *http.Request) {
	setHeaders(rw)
	switch req.Method {
	case "GET":
		fmt.Fprintf(rw, "[\"" + strings.Join(colors, "\",\"") + "\"]")
	default:
		rw.WriteHeader(405)
	}
}

func updateColors(rw http.ResponseWriter, req *http.Request) {
	setHeaders(rw)
	switch req.Method {
	case "POST":
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(req.Body)
		arr := strings.Split(buffer.String(), ";")
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])
		colors[x * 10 + y] = arr[2]
		fmt.Fprintf(rw, "[\"" + strings.Join(colors, "\",\"") + "\"]")
	default:
		rw.WriteHeader(405)
	}
}

func main() {
	colors = make([]string, 100);
	for i := 0; i < 100; i++ {
		colors[i] = "255000000"
	}
	http.HandleFunc("/get", getColors)
	http.HandleFunc("/update", updateColors)
	fmt.Printf("Listening...\nPress ^C to exit\n");
	log.Fatal(http.ListenAndServe(":8082", nil))
}