package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type snippetJSON struct {
	Title       string
	Name        string
	Trigger     string
	Description string
	Scope       string
	File        string
}

var outputPath = "./output/"
var inputPath = "./input/"

// var filePath = "./example.info"
var err = ""

// error checking helper function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// main
func main() {

	var WriteFile, err = os.Create("snip.txt")
	check(err)
	defer WriteFile.Close()

	for i, fileArray := range loopFiles() {

		fmt.Printf("// reading file %s\n", fileArray)

		file, err := ioutil.ReadFile(fileArray)
		if err != nil {
			fmt.Printf("// error while reading file %s\n", fileArray)
			fmt.Printf("File error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("// defining array of struct snippetJSON")
		var snippet []snippetJSON

		err = json.Unmarshal(file, &snippet)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

		fmt.Println("// loop over array of structs of snippetJSON")

		for i := range snippet {
			// fmt.Printf("The ship '%s' first appeared on '%s'\n", snippet[i].Name, snippet[i].Trigger)
			fmt.Printf("Title: '%s'\n", snippet[i].Title)
			fmt.Printf("Name: '%s'\n", snippet[i].Name)
			fmt.Printf("Trigger: '%s'\n", snippet[i].Trigger)
			fmt.Printf("Description: '%s'\n", snippet[i].Description)
			fmt.Printf("File: '%s'\n", snippet[i].File)

			// fmt.Fprintf(WriteFile, "snippet '%s' '%s'", snippet[i].Trigger, snippet[i].Description)
		}

		fmt.Println(i)
	}

}

func loopFiles() []string {

	fileList := []string{}
	err := filepath.Walk(inputPath, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		ignored, err := filepath.Match(".*", f.Name())
		if err != nil {
			return err
		}
		if !ignored {
			fileList = append(fileList, path)
		}
		return nil
	})
	check(err)

	// for i, file := range fileList {
	// 	fmt.Println(file)
	// }

	return fileList
}

// func VisitFile(fp string, fi os.FileInfo, err error) error {
// 	if err != nil {
// 		fmt.Println(err) // can't walk here,
// 		return nil       // but continue walking elsewhere
// 	}
// 	if fi.IsDir() {
// 		return nil // not a file.  ignore.
// 	}
// 	// ignore hidden files
// 	ignored, err := filepath.Match(".*", fi.Name())
// 	if err != nil {
// 		return err // this is fatal.
// 	}
// 	if !ignored {
// 		fmt.Println(fp)
// 	}
// 	return nil
// }
