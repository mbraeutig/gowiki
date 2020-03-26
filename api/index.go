package handler

import (
	"fmt"
	"net/http"
	"log"
)

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling url: ", r.URL.Path)
	fmt.Fprintf(w, "Url:%s", r.URL.Path)
	if r.URL.Path == "/api/edit" {
		editHandler(w, r)
	}
	if r.URL.Path == "/api/view" {
		viewHandler(w, r)
	}
	if r.URL.Path == "/api/save" {
		saveHandler(w, r)
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "editHandler")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "viewHandler")
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "saveHandler")
}
