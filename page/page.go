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

// Struct for each blog
type BlogPost struct {
    Title string
    URL   string
}

// Struct for index page
type IndexPageData struct {
	Title template.HTML
	Data  template.HTML
	Posts []BlogPost
}

// Struct for content page
type PageData struct {
	Title template.HTML
	Data  template.HTML
	

}

func GetAllPosts(contentDir string) []BlogPost {
    var posts []BlogPost

    files, err := os.ReadDir(contentDir)
    if err != nil {
        log.Fatalf("❌ cannot read content directory: %v", err)
    }

    for _, file := range files {
        if file.IsDir() || file.Name() == "_index.md" || !strings.HasSuffix(file.Name(), ".md") {
            continue
        }
        filename := strings.TrimSuffix(file.Name(), ".md")
        posts = append(posts, BlogPost{
            Title: filename,          
            URL:   filename + ".html", // link to generated HTML
        })
    }

    return posts
}


func GenarateIndexPage(contentDir string, templatePath string) {
    indexPath := contentDir + "/_index.md"

    if _, err := os.Stat(indexPath); os.IsNotExist(err) {
        log.Fatalf("❌ No _index.md found in %s", contentDir)
        return
    }

    tmpl := parse.GetIndexTemplate(templatePath)
    posts := GetAllPosts(contentDir) // get all blog posts

    data := IndexPageData{
        Title: template.HTML(GetTitle()),
        Posts: posts,
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

    fmt.Println("✔ Generated output/index.html with blog posts")
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
		Data:  template.HTML(genaratedHTML),
	}

	//Parsing filename to output the html file
	//according to it's markdown name
	_, filename := parse.GetContent(contentPath)

	//Trim the .md extension
	filename = strings.TrimSuffix(filename, ".md")

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
