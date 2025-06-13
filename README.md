# 🧰 JSON2CSV-Go

A lightweight and efficient **JSON to CSV converter** written in Go. It reads a JSON array of objects, dynamically extracts headers, and writes correct tabular CSV output—ideal for data processing workflows and automation scripts.

---

## 📋 Table of Contents

1. [Overview](#overview)  
2. [Features](#features)  
3. [Requirements](#requirements)  
4. [Installation](#installation)  
5. [Usage](#usage)  
6. [Code Structure](#code-structure)  
7. [Error Handling](#error-handling)  
8. [Contributing](#contributing)  
9. [License](#license)

---

## 💡 Overview

This Go-based CLI tool—complete with reusable package code—automatically:

- Finds the first JSON file in the working directory  
- Parses it into a slice of Go maps  
- Extracts dynamic headers from keys  
- Converts and exports as a `.csv` file  
- Supports both direct `go run` usage and `go build` for installing the binary :contentReference[oaicite:1]{index=1}

---

## ✅ Features

- 🔍 **Auto-detection**: Identifies the first `.json` file in the directory :contentReference[oaicite:2]{index=2}  
- 🧩 **Dynamic Headers**: Builds CSV headers based on JSON keys from first entry :contentReference[oaicite:3]{index=3}  
- 💾 **CSV Export**: Writes a `.csv` file with the same base name as input :contentReference[oaicite:4]{index=4}  
- 🚫 **Handles Missing Fields**: Missing keys render as empty CSV values :contentReference[oaicite:5]{index=5}  
- 🚨 **Error Reporting**: Simple messages for missing files or malformed JSON :contentReference[oaicite:6]{index=6}

---

## 🧾 Requirements

- Go **1.18+** (modules enabled)  
- No external dependencies—uses Go standard `encoding/json` and `encoding/csv`

---

## ⚙️ Installation

### Local build
```bash
git clone https://github.com/MisaghMomeniB/JSON2CSV-Go.git
cd JSON2CSV-Go/src
go build -o json2csv main.go
````

### Run directly

```bash
go run src/main.go
```

---

## 🚀 Usage

Place a JSON file (array of objects) in your folder and execute:

```bash
# e.g. data.json → data.csv
./json2csv
```

The tool reads `*.json`, converts it to CSV, and writes `*.csv`.

---

## 📁 Code Structure

```
JSON2CSV-Go/
├── src/
│   └── main.go         # CLI entrypoint and conversion logic
├── README.md           # This file
└── LICENSE             # MIT License
```

* `main.go`:

  * Finds JSON file via `filepath.Glob("*.json")`
  * Loads JSON into `[]map[string]interface{}`
  * Extracts keys from first element
  * Outputs CSV using `encoding/csv`

---

## ⚠️ Error Handling

* Exits with a message if no JSON files found
* Reports parsing errors for invalid JSON
* Gracefully handles write failures

---

## 🤝 Contributing

Improvements welcome! Consider adding:

* Support for nested JSON arrays or pointers
* CLI flags (e.g., custom file paths, headers)
* Batch file processing

To contribute:

1. Fork the repo
2. Create a feature branch
3. Submit a PR with clear descriptions

---

## 📄 License

Distributed under the **MIT License**. See [LICENSE](LICENSE) for details.
