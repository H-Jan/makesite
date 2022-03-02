package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Page struct {
	// Content of type string
	Content string
}

func main() {

	fileParse := flag.String("file", "", "The file to read in")
	flag.Parse()

	//Have file as string
	linkHTML := strings.Split(*fileParse, ".")[0] + ".html"
	//Now linkHTML will be string name of our parsed file stored in index 0 location where pointer is looking

	// use ioutil library with built in function of ReadFile to read
	fileContents, _ := ioutil.ReadFile(*fileParse)

	// Struct to display contents of our page
	page := Page{
		Content: string(fileContents),
	}

	//Create the template to render into our new HTML file
	t, _ := template.ParseFiles("template.tmpl")

	// Call in linkHTML to create the new file
	newFile, _ := os.Create(linkHTML)

	err := t.Execute(newFile, page)

	if err != nil {
		panic(err)
	}
}
