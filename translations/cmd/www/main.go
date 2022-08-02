package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

var addr string = ":4018"

func main() {
	mux := pat.New()
	mux.Get("/:locale", http.HandlerFunc(handleHome))

	log.Print("Starting server on: 4018...")
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}
