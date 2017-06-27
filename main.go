package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type snippetJSON struct {
	Title       string
	Name        string
	Trigger     string
	Description string
	Scope       string
	File        string
}

var filePath = "./example.info"
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

	fmt.Printf("// reading file %s\n", filePath)

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("// error while reading file %s\n", filePath)
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

		fmt.Fprintf(WriteFile, "snippet '%s' '%s'", snippet[i].Trigger, snippet[i].Description)

	}

}
