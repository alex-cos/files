# Files - Go File Operations Library

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue)](https://go.dev/)
[![Test Status](https://github.com/alex-cos/files/actions/workflows/test.yml/badge.svg)](https://github.com/alex-cos/files/actions/workflows/test.yml)
[![Codecov](https://codecov.io/gh/alex-cos/files/branch/main/graph/badge.svg)](https://codecov.io/gh/alex-cos/files)
[![Lint Status](https://github.com/alex-cos/files/actions/workflows/lint.yml/badge.svg)](https://github.com/alex-cos/files/actions/workflows/lint.yml)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/alex-cos/files)](https://goreportcard.com/report/github.com/alex-cos/files)

A comprehensive Go library for file operations, filtering, and file type detection.

## Features

- **File Operations**: Copy, move, delete, and concatenate files
- **Directory Operations**: Copy directories recursively, calculate directory statistics
- **File Filtering**: Flexible filters by name, extension, size, date, and regex
- **Directory Filtering**: Filter directories by name, size, and file count
- **File Type Detection**: Comprehensive file description with MIME types for 280+ extensions
- **Archive Support**: Extract and create ZIP and TAR archives
- **Hash Functions**: Calculate MD5, SHA1, SHA256, and SHA512 file hashes
- **Cross-platform**: Windows and Unix compatible

## Installation

```bash
go get github.com/alex-cos/files
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/alex-cos/files"
)

func main() {
    // Get file information
    desc, exists := files.GetFileDesc("pdf")
    if exists {
        fmt.Printf("Extension: %s, MIME: %s\n", desc.Name, desc.MimeType)
    }

    // Copy a file
    err := files.CopyFile("source.txt", "destination.txt")

    // List files with filters
    files, err := files.ListFiles(".", func(f *files.FileInfo) bool {
        return files.FilterFileByExt("go")(f)
    })
}
```

## File Operations

### Copy Files

```go
// Copy a single file
err := files.CopyFile("source.txt", "destination.txt")

// Copy to a directory (keeps filename)
err := files.CopyFile("file.txt", "/path/to/directory")

// Copy entire directory
err := files.CopyDir("/source/dir", "/destination/dir")
```

### Move Files

```go
// Move a single file
err := files.MoveFile("source.txt", "destination.txt")

// Move to a directory (keeps filename)
err := files.MoveFile("file.txt", "/path/to/directory")

// Move entire directory
err := files.MoveDir("/source/dir", "/destination/dir")
```

### Delete Files

```go
// Delete a single file
err := files.DeleteFile("file.txt")

// Delete a directory and all contents
err := files.DeleteDir("/path/to/directory")
```

### Concatenate Files

```go
// Concatenate multiple files into one
err := files.ConcatFiles([]string{"file1.txt", "file2.txt"}, "combined.txt", 0644)

// Concatenate all files of same type in a directory
err := files.ConcatDir("/source/dir", "all.txt", files.FilterFileByExt("log"), 0644)
```

### Hash Files

```go
// Calculate MD5 hash
hash, err := files.FileHash("file.txt", files.MD5)

// Calculate SHA256 hash
hash, err := files.FileHash("file.txt", files.SHA256)
```

### Directory Statistics

```go
stats, err := files.GetDirStats("/path/to/dir")
fmt.Printf("Total files: %d, Total size: %s\n", stats.TotalFiles, stats.FormatSize())
```

## Filtering

### File Filters

```go
// Filter by extension
filter := files.FilterFileByExt("go")

// Filter by name
filter := files.FilterFileByName("README.md")

// Filter by size (greater than)
filter := files.FilterFileBySizeGreater(1024 * 1024)

// Filter by date (updated after)
filter := files.FilterFileByUpdatedAfter(time.Now().AddDate(0, 0, -7))

// Filter by regex
filter := files.FilterFileByRegEx(*regexp.MustCompile(`^test.*\.txt$`))

// Filter by category
filter := files.FilterFileByCategory(files.CategoryCode)

// Filter by MIME type
filter := files.FilterFileByMimeType(files.MimeJSON)

// Combine filters
combined := func(f *files.FileInfo) bool {
    return files.FilterFileByExt("go")(f) && files.FilterFileBySizeGreater(100)(f)
}
```

### Directory Filters

```go
// Filter by name
filter := files.FilterDirByName("node_modules")

// Filter by size
filter := files.FilterDirBySizeGreater(1024 * 1024 * 100)

// Filter by file count
filter := files.FilterDirByNbFilesGreater(10)
```

## Listing Files

```go
// List all files in a directory
files, err := files.ListFiles("/path/to/dir", nil)

// List files with filter
txtFiles, err := files.ListFiles("/path/to/dir", files.FilterFileByExt("txt"))

// List directories
dirs, err := files.ListDirs("/path/to/parent", nil)
```

## Archive Operations

### ZIP Archives

```go
// Create a ZIP archive
err := files.ZipAll("/source/dir", "/output/archive.zip")

// Extract a ZIP archive
err := files.UnZip("/archive.zip", "/output/dir")

// List files in a ZIP archive
names, err := files.ZipList("/archive.zip")
```

### TAR Archives

```go
// Create a TAR archive
err := files.TarAll("/source/dir", "/output/archive.tar")

// Extract a TAR archive
err := files.UnTar("/archive.tar", "/output/dir")

// List files in a TAR archive
names, err := files.TarList("/archive.tar")
```

## File Descriptions

Get detailed information about file types:

```go
// Get file description
desc, exists := files.GetFileDesc("pdf")
if exists {
    fmt.Printf("Name: %s\n", desc.Name)
    fmt.Printf("Category: %s\n", desc.Category)
    fmt.Printf("MIME Type: %s\n", desc.MimeType)
    fmt.Printf("Binary: %v\n", desc.IsBinary)
    fmt.Printf("Compressed: %v\n", desc.IsCompressed)
}

// Get file name by extension
name := files.GetFileDescName("jpg") // "JPEG Image"

// Get file category
cat := files.GetFileDescCat("go") // "Code"

// Get MIME type
mime := files.GetMimeType("json") // "application/json"

// Check file properties
isBinary := files.IsFileBinary("pdf")   // *bool
isCompressed := files.IsFileCompressed("zip") // *bool
```

## Constants

### Hash Algorithms

```go
files.MD5
files.SHA1
files.SHA256
files.SHA512
```

### File Categories

```go
files.CategoryDocument
files.CategorySpreadsheet
files.CategoryPresentation
files.CategoryImage
files.CategoryAudio
files.CategoryVideo
files.CategoryArchive
files.CategoryExecutable
files.CategoryCode
files.CategoryConfig
files.CategoryData
files.CategoryDatabase
files.CategorySecurity
files.CategoryFont
files.Category3D
files.CategoryEbook
files.CategoryCloud
files.CategoryAutomation
files.CategoryScientist
files.CategoryDiskImage
files.CategoryLog
files.CategoryTemp
files.CategoryOther
```

## Data Types

### DirInfo

```go
type DirInfo struct {
    Path    string
    Name    string
    NbDirs  int64
    Nbfiles int64
    Size    int64
}
```

### FileInfo

```go
type FileInfo struct {
    Path         string
    Name         string
    Ext          string
    Size         int64
    Created      time.Time
    Updated      time.Time
    IsExecutable bool
}
```

### DirStats

```go
type DirStats struct {
    TotalFiles  int64
    TotalDirs   int64
    TotalSize   int64
    OldestFile  time.Time
    NewestFile  time.Time
    AverageSize int64
}
```

### FileDesc

```go
type FileDesc struct {
    Name         string
    Category     string
    MimeType     string
    IsBinary     bool
    IsCompressed bool
    IsExecutable bool
}
```
