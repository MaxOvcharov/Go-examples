package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const rootPath string = "Simple_web_app/"
const templatePathPrefix string = "templates/"
const publishPathPrefix string = "publish/"

const urlEditPath string = "/edit/"
const urlSavePath string = "/save/"
const urlViewPath string = "/view/"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := rootPath + publishPathPrefix + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(w http.ResponseWriter, r *http.Request, urlPath string) (*Page, error) {
	title := r.URL.Path[len(urlPath):]
	filename := rootPath + publishPathPrefix + title + ".txt"

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		http.Redirect(w, r, urlEditPath + title, http.StatusNotFound)
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, p *Page, tmpl string) {
	tmplPath := rootPath + templatePathPrefix + tmpl
	t, _ := template.ParseFiles(tmplPath)
	t.Execute(w, p)
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage(w, r, urlViewPath)
	renderTemplate(w, p, "view.html")

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(w, r, urlEditPath)

	if err != nil {
		title := r.URL.Path[len(urlEditPath):]
		p = &Page{Title: title}
	}

	renderTemplate(w, p, "edit.html")
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(urlSavePath):]
	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}
	p.save()

	http.Redirect(w, r, urlViewPath + title, http.StatusFound)
}

func main() {
	http.HandleFunc("/", simpleHandler)
	http.HandleFunc(urlViewPath, viewHandler)
	http.HandleFunc(urlEditPath, editHandler)
	http.HandleFunc(urlSavePath, saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}