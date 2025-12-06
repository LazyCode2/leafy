package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/LazyCode2/leafy/config"
	"github.com/LazyCode2/leafy/page"
)

func main() {
	buildFlag := flag.Bool("build", false, "Build your site")
	initFlag := flag.Bool("init", false, "Initialize project structure")
	flag.Parse()

	cfg , err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println("Failed to load config")
	}

	if *initFlag {
		BuildNewStructure()
	} else if *buildFlag {
	    fmt.Println("Building site...")
	    page.GenarateIndexPage(cfg.ContentDir, cfg.TemplatePath)

	    files, err := os.ReadDir(cfg.ContentDir)
	    if err != nil {
	        fmt.Println("Failed to read content directory:", err)
	        return
	    }

	    for _, f := range files {
	        if f.IsDir() {
	            continue // Skip directories
	        }
	        name := f.Name()
	        if name == "_index.md" || filepath.Ext(name) != ".md" {
	            continue // Skip _index.md and non-markdown files
	        }
	        // Generate a page for each markdown file
	        page.GenaratePage(filepath.Join(cfg.ContentDir, name), cfg.TemplatePath)
		}
	} else {
		fmt.Println("No flag provided. \nUse leafy --build or --init")
	}

}

func BuildNewStructure() {
	dirs := []string{
		"content",
		"output",
		"template",
	}

	// Creating directories
	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory", dir, ":", err)
			continue
		}
		fmt.Println("Created directory:", dir)
	}

	// Creating default config.yaml
	configContent := `title: "Hello world"
template_path: "template/default.tmpl"
content_dir: "content/"
`
	configPath := "config.yaml"
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		fmt.Println("Error creating config.yaml:", err)
	} else {
		fmt.Println("Created config.yaml")
	}

	// Creating default template
	templateContent := `<!DOCTYPE html>
<html>
<head>
    <title>{{ .Title }}</title>
</head>
<body>
    {{ .Data }}
</body>
</html>
`
	templateFilePath := filepath.Join("template", "default.tmpl")
	err = os.WriteFile(templateFilePath, []byte(templateContent), 0644)
	if err != nil {
		fmt.Println("Error creating default template:", err)
	} else {
		fmt.Println("Created default template:", templateFilePath)
	}

	fmt.Println("Project structure initialized!")
}

