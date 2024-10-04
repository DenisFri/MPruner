# MPruner

## Introduction
MPruner is a script designed to clean up cached files. It allows you to specify a target directory and choose whether to delete all files or only the last modified file in the directory. The operation results, including any errors encountered, are logged to a log file. Multiple directories can be configured for cleaning.

## Features

- **Configurable Path**: Define the target directory through a configuration file.
- **Multiple Directories**: Configure multiple directories for cleaning.
- **Selective Deletion**: Choose to delete all files or only the last modified file in the directory.
- **Logging**: Logs the operation results, including any errors encountered, to a log file.

## Requirements

- Go 1.23.0 or later

## Setup

### 1. Clone the Repository

Clone the repository to your local machine.

```bash
git clone https://github.com/DenisFri/MissionCachePruner.git
cd mpruner
```

### 2. Configuration

In case the default configuration file is missing, create a new `config.json` by copying the provided template.

```json
{
  "directories": [
    {
      "path": "Mention\\Your\\Directory\\Here",
      "delete_all": false
    },
    {
      "path": "Mention\\Your\\Second\\Directory\\Here",
      "delete_all": false
    }
  ]
}
```

### 3. Build

To build the Go script into an executable, run:

```go
go build -o mpruner.exe main.go
```

### 4. Run

To run the program, use:

```bash
./mpruner.exe
```

## Logging

MPruner logs all its operations to `cleanup.log`. The log file records:

- Expanded directory paths.
- Successful or failed attempts to delete files.
- Information about skipping duplicate directories.
- Any errors encountered during the process.

