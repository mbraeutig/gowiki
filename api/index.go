package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url:%s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
