package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Hello Go!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
