package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Page struct {
	// Content of type string
	Content string
}

func main() {

	//NOTE: COMMENTING OUT PREVIOUS WORK FOR EASE OF WORKING ON NEW VERSIONS
	//fileParse := flag.String("file", "latest-post.txt", "The file to read in")

	//V1.1 Flag for input directory
	dDir := flag.String("dir", "", "The input directory")
	flag.Parse()

	// Looped statement
	if *dDir != "" {
		// files in directory
		files, err := ioutil.ReadDir((*dDir))
		if err != nil {
			panic(err)
		}
		// loop through each file throughout all our files
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".txt" {
				// This is essentially the same previous code as V 1.0
				linked_HTML := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + ".html"
				filed_Contents, _ := ioutil.ReadFile(file.Name())

				struct_new := Page{Content: string(filed_Contents)}

				new_t, _ := template.ParseFiles("template.tmpl")

				newFile_two, _ := os.Create(linked_HTML)

				err := new_t.Execute(newFile_two, struct_new)
				if err != nil {
					panic(err)
				}
			}

		}
	}

	//NOTE: COMMENTING OUT PREVIOUS WORK FOR EASE OF WORKING ON NEW VERSIONS
	// THE ABOVE LOOP IS ESSENTIALLY THE SAME CODE BUT LOOPED FOR DIRECTORY
	//Have file as string
	//linkHTML := strings.Split(*fileParse, ".")[0] + ".html"
	//Now linkHTML will be string name of our parsed file stored in index 0 location where pointer is looking

	// use ioutil library with built in function of ReadFile to read
	//fileContents, _ := ioutil.ReadFile(*fileParse)

	// Struct to display contents of our page
	//page := Page{
	//	Content: string(fileContents),
	//}

	//Create the template to render into our new HTML file
	//t, _ := template.ParseFiles("template.tmpl")

	// Call in linkHTML to create the new file
	//newFile, _ := os.Create(linkHTML)

	//err := t.Execute(newFile, page)

	//if err != nil {
	//	panic(err)
	//}
}
