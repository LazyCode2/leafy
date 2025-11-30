package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type PageData struct {
	Data template.HTML // must be exported for templates
}

func GenaratePage(){
	tmpl    := GetTemplate("template/default.tmpl")
	genaratedHTML := GenarateHTMLbody()
	data := PageData{
	 Data: template.HTML(genaratedHTML),
	}
	
	outFile, err := os.Create("output/index.html")
	if err != nil {
		log.Fatalf("cannot create output file: %v", err)
	}
	defer outFile.Close()
	err = tmpl.Execute(outFile, data)
	
	if err != nil {
		log.Fatalf("template execute failed: %v", err)
	}
	fmt.Println("✔ Generated output/index.html")
}