package page

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/LazyCode2/leafy/config"
	"github.com/LazyCode2/leafy/genarate"
	"github.com/LazyCode2/leafy/parse"
)

type PageData struct {
	Title template.HTML
	Data  template.HTML // must be exported for templates

}

func GenarateIndexPage(contentDir string, templatePath string) {
    indexPath := contentDir + "/_index.md"

    if _, err := os.Stat(indexPath); os.IsNotExist(err) {
        log.Fatalf("❌ No _index.md found in %s", contentDir)
        return
    }

   
    tmpl := parse.GetTemplate(templatePath)  // load template
    genaratedHTML := genarate.GenarateHTMLbody(indexPath)

    data := PageData{
        Title: template.HTML(GetTitle()),
        Data:  template.HTML(genaratedHTML),
    }

    os.MkdirAll("output", os.ModePerm)
    outFile, err := os.Create("output/index.html")
    if err != nil {
        log.Fatalf("❌ cannot create output file: %v", err)
    }
    defer outFile.Close()

    err = tmpl.Execute(outFile, data)
    if err != nil {
        log.Fatalf("❌ template execute failed: %v", err)
    }

    fmt.Println("✔ Generated output/index.html from _index.md")
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

	//Parsing filename to output the html file
	//according to it's markdown name
	_ , filename := parse.GetContent(contentPath)

	//Trim the .md extension
	filename = strings.TrimSuffix(filename,".md") 

	outFile, err := os.Create("output/" + filename + ".html")
	if err != nil {
	    log.Fatalf("cannot create output file: %v", err)
	}
	defer outFile.Close()
	err = tmpl.Execute(outFile, data)

	if err != nil {
		log.Fatalf("template execute failed: %v", err)
	}
	fmt.Println("✔ Generated output/" + filename + ".html")
}
