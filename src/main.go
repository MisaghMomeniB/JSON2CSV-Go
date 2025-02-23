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