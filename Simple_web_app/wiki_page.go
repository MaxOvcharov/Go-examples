package main

import (
	"fmt"
	"io/ioutil"
)

type PageV1 struct {
	Title string
	Body []byte
}

func (p *PageV1) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPageV1(title string) (*PageV1, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	
	return &PageV1{Title: title, Body: body}, nil
}

func main()  {
	title := "TestWikiPageV1"
	p1 := &PageV1{Title: title, Body: []byte("This is test Page!")}
	p1.save()
	p2, _ := loadPageV1(title)
	fmt.Println(string(p2.Body))
}