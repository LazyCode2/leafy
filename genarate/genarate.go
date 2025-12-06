package genarate

import (
	"log"

	"github.com/LazyCode2/leafy/parse"
	"github.com/gomarkdown/markdown"
)

func GenarateHTMLbody(contentPath string) []byte {
	content , _ := parse.GetContent(contentPath)

	generatedHTML := markdown.ToHTML(content, nil, nil)
	if generatedHTML == nil{
		log.Fatalf("Failed to genarate HTMLbody")
	}

	return generatedHTML
}
