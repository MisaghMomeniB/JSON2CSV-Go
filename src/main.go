package main

import (
	"encoding/csv"  // Import package for CSV file handling
	"encoding/json" // Import package for JSON parsing
	"fmt"           // Import package for formatted I/O
	"io/ioutil"     // Import package for file reading
	"os"            // Import package for file operations
	"path/filepath" // Import package for working with file paths
)

func main() {
	// Find all JSON files in the current directory
	jsonFiles, err := filepath.Glob("*.json")
	if err != nil || len(jsonFiles) == 0 {
		fmt.Println("No JSON files found.") // Print message if no JSON files exist
		return                              // Exit the program
	}

	// Use the first JSON file found
	jsonFileName := jsonFiles[0]
	fmt.Println("Processing file:", jsonFileName) // Print the name of the file being processed

	// Read the JSON file
	jsonFile, err := ioutil.ReadFile(jsonFileName)
	if err != nil {
		fmt.Println("Error reading JSON file:", err) // Print error if file cannot be read
		return                                      // Exit the program
	}

	// Convert JSON to a slice of maps
	var records []map[string]interface{}
	err = json.Unmarshal(jsonFile, &records)
	if err != nil {
		fmt.Println("Error processing JSON:", err) // Print error if JSON parsing fails
		return                                    // Exit the program
	}

	// Check if JSON contains data
	if len(records) == 0 {
		fmt.Println("JSON file is empty or has an invalid format.") // Print message if JSON is empty
		return                                                     // Exit the program
	}

	// Create CSV file (same name as JSON but with .csv extension)
	csvFileName := jsonFileName[:len(jsonFileName)-5] + ".csv" // Change file extension from .json to .csv
	csvFile, err := os.Create(csvFileName)
	if err != nil {
		fmt.Println("Error creating CSV file:", err) // Print error if file creation fails
		return                                      // Exit the program
	}
	defer csvFile.Close() // Ensure the file is closed when the function exits

	writer := csv.NewWriter(csvFile) // Create a new CSV writer
	defer writer.Flush()              // Ensure data is written to file before closing

	// Write CSV headers (JSON keys)
	headers := make([]string, 0, len(records[0])) // Create a slice for CSV headers
	for key := range records[0] {                 // Extract keys from the first JSON object
		headers = append(headers, key) // Add each key as a CSV header
	}
	writer.Write(headers) // Write headers to the CSV file

	// Write data rows
	for _, record := range records {         // Loop through each JSON object
		row := make([]string, len(headers)) // Create a slice for CSV row data
		for i, header := range headers {    // Loop through headers
			if val, ok := record[header]; ok {
				row[i] = fmt.Sprintf("%v", val) // Convert value to string format
			} else {
				row[i] = "" // Leave empty if key is missing in this JSON object
			}
		}
		writer.Write(row) // Write row data to CSV file
	}

	fmt.Println("JSON to CSV conversion successful! Output:", csvFileName) // Print success message
}
