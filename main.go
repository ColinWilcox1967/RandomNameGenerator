package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// Data represents the expected structure of the JSON input
type Data struct {
	Items []string `json:"names"`
}

func main() {

	// very simplistic error handling for arguments
	if len(os.Args) < 2 {
		fmt.Println("Error: No input file specified.")
		os.Exit(-1)
	}

	if len(os.Args) > 2 {
		fmt.Println ("Error: Too many arguments specified.")
		os.Exit(-2)
	}

	// Only interested in the first argument (after the .EXE name) for the input data file
	filePath := os.Args[1]

	// Now try and read whats in it
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: Could not read file '%s': %v\n", filePath, err)
		os.Exit(-3)
	}

	// Get the name list from the input file
	var data Data
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		fmt.Printf("Error: Failed to parse JSON from file '%s': %v\n", filePath, err)
		os.Exit(-4)
	}

	// "names" tag is defined but theres nothing inside
	if len(data.Items) == 0 {
		fmt.Println("Error: No names found in the 'names' field.")
		os.Exit(1)
	}

	fmt.Printf("Total names loaded is %d\n", len(data.Items))

	// seed RNG with current time and choose one
	rand.Seed(time.Now().UnixNano())
	selected := data.Items[rand.Intn(len(data.Items))]

	fmt.Printf("Randomly selected name is *** %s ***\n", selected)
}
