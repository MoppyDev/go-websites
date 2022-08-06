package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {

	websiteName := flag.String("websiteName", "default-go-website", "The website name")

	if len(os.Args) <= 1 {
		fmt.Printf("No website name passed. Using default value [%v]\n", *websiteName)
	} else if len(os.Args) > 2 {
		fmt.Print("Too many arguments passed.\n")
		return
	}

	flag.Parse()
	fmt.Printf("Creating recommended folders for your new Go website [%v]\n", *websiteName)

	makeDir(*websiteName, 0755)
	makeDir(*websiteName+"/internal", 0755)
	makeDir(*websiteName+"/cmd/web", 0755)
	makeDir(*websiteName+"/ui/html/pages", 0755)
	makeDir(*websiteName+"/ui/html/partials", 0755)
	makeDir(*websiteName+"/ui/static/css", 0755)
	makeDir(*websiteName+"/ui/static/img", 0755)
	makeDir(*websiteName+"/ui/static/js", 0755)

	makeEmptyFile(*websiteName+"/cmd/web/main.go", 0755)
	makeEmptyFile(*websiteName+"/cmd/web/handlers.go", 0755)
}

func makeDir(path string, permission fs.FileMode) {
	fmt.Printf("Creating path: %v\n", path)
	err := os.MkdirAll(path, permission)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

func makeEmptyFile(path string, permission fs.FileMode) {
	err := os.WriteFile(path, nil, permission)
	if err != nil {
		log.Fatal(err)
	}
}
