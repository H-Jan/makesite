package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

type Page struct {
	//HTMLPagePath and Content of type string
	HTMLPagePath string
	Content      string
}

func main() {
	// use ioutil library with built in function of ReadFile to read
	// create error statement to abort program if problem encountered
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	page := Page{
		HTMLPagePath: "first-post.html",
		Content:      string(fileContents),
	}

	//Create the template to render into our new HTML file
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// Call in HTMLPagePath to show what page to create
	newFile, err := os.Create(page.HTMLPagePath)
	if err != nil {
		panic(err)
	}

	t.Execute(newFile, page)
}
