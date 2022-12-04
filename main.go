package main

import (
	"log"
	"net/http"

	"github.com/mrayone/learn-go/di"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
