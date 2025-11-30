package main

import (
	"html/template"
	"log"
	"os"
)

// Reads a markdown file from the given path and returns its raw bytes.
func GetContent(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("cannot read markdown: %v", err)
	}

	return content
}

// Loads and parses the main HTML template used for page generation.
func GetTemplate(path string) *template.Template {
	tpl, err := template.ParseFiles("template/default.tmpl")
	if err != nil {
		log.Fatalf("cannot load template: %v", err)
	}

	return tpl
}