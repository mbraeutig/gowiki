package handler

import (
	"fmt"
	"net/http"
)

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {
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

func init() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
