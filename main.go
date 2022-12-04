package main

import (
	"learn-go-with-tests/di"
	"log"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
