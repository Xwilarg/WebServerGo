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

func getRequest(rw http.ResponseWriter, http *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Content-Type", "application/json")
	switch http.Method {
	case "GET":
		fmt.Fprintf(rw, "[\"" + strings.Join(colors, "\",\"") + "\"]")
	case "POST":
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(http.Body)
		arr := strings.Split(buffer.String(), ";")
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])
		colors[x * 10 + y] = arr[2]
		fmt.Fprintf(rw, "[\"" + strings.Join(colors, "\",\"") + "\"]")
	}
}

func main() {
	colors = make([]string, 100);
	for i := 0; i < 100; i++ {
		colors[i] = "255000000"
	}
	http.HandleFunc("/", getRequest)
	fmt.Printf("Listening...\nPress ^C to exit\n");
	log.Fatal(http.ListenAndServe(":8082", nil))
}