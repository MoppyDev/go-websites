package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	websiteName := flag.String("websiteName", "default-go-website", "The website name")

	if len(os.Args) <= 1 {
		fmt.Printf("No website name passed. Using default value [%v]\n", *websiteName)
	} else if len(os.Args) == 2 {
		log.Fatal("Utilize -websiteName flag and then pass parameter\n")
	} else if len(os.Args) > 3 {
		log.Fatal("Too many arguments passed.\n")
	}

	flag.Parse()

	if checkIfDirExists(*websiteName) {
		log.Printf("Exiting.. Folder already exists: [%v]\n", *websiteName)
		return
	} else {
		fmt.Println("Folder doesn't yet exist.")
	}

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

func checkIfDirExists(dir string) bool {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Searching for %v\n", dir)

	for _, file := range files {
		if file.Name() == dir {
			fmt.Printf("Folder already exists [%v]\n", dir)
			return true
		}
	}

	return false
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
