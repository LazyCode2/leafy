package parse

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

// Reads a markdown file from the given path and returns its raw bytes and filename.
func GetContent(path string) ([]byte, string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("cannot read markdown: %v", err)
	}

	filename := filepath.Base(path)

	return content, filename
}

// Loads and parses the main HTML template used for page generation.
func GetTemplate(path string) *template.Template {
	tpl, err := template.ParseFiles("template/default.tmpl")
	if err != nil {
		log.Fatalf("cannot load template: %v", err)
	}

	return tpl
}

// Load and Parse the index HTML template for index page
func GetIndexTemplate(path string) *template.Template {
	tpl, err := template.ParseFiles("template/index.tmpl")
	if err != nil {
		log.Fatalf("cannot load template: %v", err)
	}

	return tpl
}
