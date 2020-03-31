package handler

import (
	"fmt"
	"io/ioutil"
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

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func init() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
