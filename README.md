# 📜 JSON to CSV Converter in Go  

## 🚀 Overview  
This Go program reads a JSON file from the current directory, converts its data into CSV format, and saves it as a new file. It automatically detects JSON files and processes the first one found. The program is simple, efficient, and useful for transforming structured JSON data into a tabular CSV format.

---

## 🔧 Features  
✅ **Automatic File Detection** – Finds the first JSON file in the directory.  
✅ **JSON Parsing** – Reads and decodes JSON into a structured format.  
✅ **Dynamic CSV Generation** – Extracts headers dynamically from JSON keys.  
✅ **Error Handling** – Provides clear messages for errors like missing files or malformed JSON.  
✅ **Efficient Writing** – Uses Go’s built-in CSV writer for optimal performance.  

---

## 📂 File Conversion Process  
1️⃣ **Finds a JSON file** in the current directory.  
2️⃣ **Reads the file** and parses its contents into a Go slice of maps.  
3️⃣ **Extracts keys** from the JSON data to generate CSV headers.  
4️⃣ **Writes the CSV file**, converting values to strings.  
5️⃣ **Saves the output** with the same filename but with a `.csv` extension.  

---

## 🛠️ Installation & Usage  

### 🔹 Prerequisites  
Make sure you have **Go installed** on your system. If not, download it from [golang.org](https://golang.org/dl/).  

### 🔹 Running the Program  

1️⃣ Clone this repository or copy the `main.go` file to your working directory.  
```bash
git clone https://github.com/yourusername/json-to-csv-go.git
cd json-to-csv-go
```
   
2️⃣ Place a JSON file in the same directory as `main.go`. Example JSON format:
```json
[
    {"name": "Alice", "age": 25, "city": "New York"},
    {"name": "Bob", "age": 30, "city": "San Francisco"}
]
```

3️⃣ Run the program:  
```bash
go run main.go
```

4️⃣ The CSV file will be generated in the same directory with the same name as the JSON file.  

---

## 📝 Example Output  

### ✅ Input (`data.json`):  
```json
[
    {"name": "John", "age": 28, "country": "USA"},
    {"name": "Sara", "age": 24, "country": "Canada"}
]
```

### ✅ Output (`data.csv`):  
```csv
name,age,country
John,28,USA
Sara,24,Canada
```

---

## 🏗 Code Explanation  

### 1️⃣ **Finding JSON Files**  
The program uses `filepath.Glob("*.json")` to find JSON files in the current directory. If no files are found, it prints a message and exits.  

### 2️⃣ **Reading & Parsing JSON**  
The JSON file is read using `ioutil.ReadFile()` and then parsed into a slice of maps using `json.Unmarshal()`. Each map represents a JSON object.  

### 3️⃣ **Extracting Headers for CSV**  
The first JSON object is used to extract keys dynamically, which become the headers for the CSV file.  

### 4️⃣ **Writing Data to CSV**  
A new CSV file is created, and the extracted headers are written as the first row. Each JSON object is then written as a new row, with values converted to strings using `fmt.Sprintf("%v", val)`.  

### 5️⃣ **Error Handling**  
Errors like missing files, unreadable JSON, or file creation failures are handled gracefully with informative messages.  

---

## 📌 Notes  
⚡ **Only the first JSON file found is processed.**  
⚡ **Assumes JSON is an array of objects (not a single object).**  
⚡ **Handles missing keys by inserting empty values in CSV.**  

---

## 📜 License  
This project is open-source and licensed under the MIT License. Feel free to use, modify, and share! 🚀  

---

## 💡 Contributions  
Have an idea to improve this tool? Feel free to fork, submit issues, or create a pull request. Contributions are always welcome! 🎉  

Happy coding! 💻🚀
