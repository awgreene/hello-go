package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/env", env)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

// says hello to world environment variable.
func env(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", os.Getenv("WORLD"))
}
