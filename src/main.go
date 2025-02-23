package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Read JSON file
	jsonFile, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Convert JSON to a slice of maps
	var records []map[string]interface{}
	err = json.Unmarshal(jsonFile, &records)
	if err != nil {
		fmt.Println("Error processing JSON:", err)
		return
	}

	// Check if JSON contains data
	if len(records) == 0 {
		fmt.Println("JSON file is empty or has an invalid format.")
		return
	}