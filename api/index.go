package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page ...
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "/tmp/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "/tmp/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/view/" || r.URL.Path == "/api/view" {
		viewHandler(w, r)
	}
	if r.URL.Path == "/api/edit/" || r.URL.Path == "/api/edit" {
		editHandler(w, r)
	}
	if r.URL.Path == "/api/save/" || r.URL.Path == "/api/save" {
		saveHandler(w, r)
	}
	if r.URL.Path == "/api/test/" || r.URL.Path == "/api/test" {
		testHandler(w, r)
	}
	if r.URL.Path == "/api/ListFiles/" || r.URL.Path == "/api/ListFiles" {
		ListFiles(w, r)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>testHandler</h1>")
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	err := p1.save()
	if err != nil {
		fmt.Fprintf(w, "<h1>Error: %s</h1>", err)
		fmt.Printf("Save error: %s", err)
		return
	}

	p2, err := loadPage("TestPage")
	if err != nil {
		fmt.Printf("load error: %s", err)
		return
	}
	fmt.Println(string(p2.Body))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func init() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
}

// ListFiles lists the files in the current directory.
func ListFiles(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("../")
	if err != nil {
		http.Error(w, "Unable to read files", http.StatusInternalServerError)
		log.Printf("ioutil.ListFiles: %v", err)
		return
	}
	fmt.Fprintln(w, "Files:")
	for _, f := range files {
		fmt.Fprintf(w, "\t%v\n", f.Name())
	}
}
