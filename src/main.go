package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// Find all JSON files in the current directory
	jsonFiles, err := filepath.Glob("*.json")
	if err != nil || len(jsonFiles) == 0 {
		fmt.Println("No JSON files found.")
		return
	}

	// Use the first JSON file found
	jsonFileName := jsonFiles[0]
	fmt.Println("Processing file:", jsonFileName)

	// Read the JSON file
	jsonFile, err := ioutil.ReadFile(jsonFileName)
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

	// Create CSV file (same name as JSON but with .csv extension)
	csvFileName := jsonFileName[:len(jsonFileName)-5] + ".csv"
	csvFile, err := os.Create(csvFileName)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write CSV headers (JSON keys)
	headers := make([]string, 0, len(records[0]))
	for key := range records[0] {
		headers = append(headers, key)
	}
	writer.Write(headers)

	// Write data rows
	for _, record := range records {
		row := make([]string, len(headers))
		for i, header := range headers {
			if val, ok := record[header]; ok {
				row[i] = fmt.Sprintf("%v", val) // Convert value to string
			} else {
				row[i] = "" // Leave empty if key is missing
			}
		}
		writer.Write(row)
	}

	fmt.Println("JSON to CSV conversion successful! Output:", csvFileName)
}
