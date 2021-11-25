package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Print("HTTP server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
