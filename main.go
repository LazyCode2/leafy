package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	buildFlag := flag.Bool("build", false, "Build your site")
	initFlag := flag.Bool("init", false, "Initialize project structure")
	flag.Parse()

	if *initFlag {
		BuildNewStructure()
	} else if *buildFlag {
		fmt.Println("Building site...")
		GenaratePage()
	} else {
		fmt.Println("No flag provided. use leaf -h")
	}
}

func BuildNewStructure(){
	dirs := []string{
		"content",
		"output",
		"template",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory", dir, ":", err)
			continue
		}
		fmt.Println("Created directory:", dir)
	}

	fmt.Println("Project structure initialized!")
}
