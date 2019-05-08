package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const rootPath string = "Simple_web_app/"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := rootPath + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := rootPath + title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, p *Page, tmpl string) {
	tmplPath := rootPath + tmpl
	t, _ := template.ParseFiles(tmplPath)
	t.Execute(w, p)
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, p, "view.htlm")

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, p, "edit.html")
}

func main() {
	http.HandleFunc("/", simpleHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}