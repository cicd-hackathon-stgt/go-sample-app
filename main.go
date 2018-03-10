package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ci-cd-hackathon/go-sample-app/greeting"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, greeting.GetGreeting())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
