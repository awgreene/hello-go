package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

// says hello to world environment variable.
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s\n", os.Getenv("WORLD"))
}
