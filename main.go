package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, GetGreeting())
}

func GetGreeting() string {
	return "Hi from Go!"
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
