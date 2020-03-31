package api

import (
	"fmt"
	"log"
	"net/http"
)


type IndexHandler struct {
	Value string
}

// Handler ...
func (i *IndexHandler) Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url:%s", r.URL.Path)

	log.Println("Handling url: ", r.URL.Path)

	if r.URL.Path == "/api/view/" {
		viewHandler(w, r)
	}
	if r.URL.Path == "/api/edit/" {
		editHandler(w, r)
	}
	if r.URL.Path == "/api/save/" {
		saveHandler(w, r)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "viewHandler")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "editHandler")
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "saveHandler")
}
