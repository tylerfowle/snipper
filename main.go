package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"
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

	fileList := []string{}
	err := filepath.Walk(inputPath, func(path string, file os.FileInfo, err error) error {
		if file.IsDir() {
			return nil
		}
		matched, err := filepath.Match("*.info", file.Name())
		if err != nil {
			return err
		}
		if matched {
			fileList = append(fileList, path)

			var WriteFile, err = os.Create("output/" + file.Name())
			check(err)
			defer WriteFile.Close()

			filejson, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Printf("// error while reading file %s\n", path)
				fmt.Printf("File error: %v\n", err)
				os.Exit(1)
			}

			var snippet []snippetJSON

			err = json.Unmarshal(filejson, &snippet)
			if err != nil {
				fmt.Println("error:", err)
				os.Exit(1)
			}

			for i := range snippet {

				fmt.Printf("Filename: ")
				fmt.Printf(file.Name() + "\n")

				fmt.Printf("Path: ")
				fmt.Printf(path + "\n")

				fmt.Printf("\n")

				fmt.Printf("Title: ")
				color.Green("'%s'\n", snippet[i].Title)

				fmt.Printf("Name: ")
				color.Green("'%s'\n", snippet[i].Name)

				fmt.Printf("Trigger: ")
				color.Green("'%s'\n", snippet[i].Trigger)

				fmt.Printf("Description: ")
				color.Green("'%s'\n", snippet[i].Description)

				fmt.Printf("File: ")
				color.Green("'%s'\n", snippet[i].File)

				fmt.Printf("\n\n")

				// fmt.Fprintf(WriteFile, "snippet '%s' '%s'", snippet[i].Trigger, snippet[i].Description)
			}

		}
		return nil
	})
	check(err)

}
