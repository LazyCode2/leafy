package main

import (
	"log"

	"github.com/gomarkdown/markdown"
)

func GenarateHTMLbody() []byte {
	content := GetContent("content/content.md")

	generatedHTML := markdown.ToHTML(content, nil, nil)
	if generatedHTML == nil{
		log.Fatalf("Failed to genarate htmlbody")
	}

	return generatedHTML
}