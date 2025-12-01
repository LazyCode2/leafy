package page

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/LazyCode2/leafy/config"
	"github.com/LazyCode2/leafy/genarate"
	"github.com/LazyCode2/leafy/parse"
)

type PageData struct {
	Title template.HTML
	Data  template.HTML // must be exported for templates

}

func GetTitle() string {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Print(err)
	}

	return cfg.Title
}

func GenaratePage(contentPath string, templatePath string) {
	tmpl := parse.GetTemplate(templatePath)
	genaratedHTML := genarate.GenarateHTMLbody(contentPath)
	data := PageData{
		Title: template.HTML(GetTitle()),
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
