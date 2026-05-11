package files

import (
	"fmt"
	"strings"
	"time"
)

// DirInfo represents information about a directory.
type DirInfo struct {
	Path    string
	Name    string
	Nbfiles int64
	Size    int64
}

// FilterDir is a function type for filtering directories.
type FilterDir func(f *DirInfo) bool

// String returns a string representation of the DirInfo.
func (item *DirInfo) String() string {
	return fmt.Sprintf("%+v", *item)
}

// FormatSize returns a human-readable string representation of the directory size.
func (item *DirInfo) FormatSize() string {
	return FormatSize(item.Size)
}

// FileInfo represents information about a file.
type FileInfo struct {
	Path    string
	Name    string
	Ext     string
	Size    int64
	Created time.Time
	Updated time.Time
}

// String returns a string representation of the FileInfo.
func (item *FileInfo) String() string {
	return fmt.Sprintf("%+v", *item)
}

// FormatSize returns a human-readable string representation of the file size.
func (item *FileInfo) FormatSize() string {
	return FormatSize(item.Size)
}

// GetExt returns the file extension in lowercase without the leading dot.
func (item *FileInfo) GetExt() string {
	return strings.ToLower(strings.TrimLeft(item.Ext, "."))
}

// FilterFile is a function type for filtering files.
type FilterFile func(f *FileInfo) bool

// DirStats represents statistics about a directory.
type DirStats struct {
	TotalFiles  int64
	TotalDirs   int64
	TotalSize   int64
	OldestFile  time.Time
	NewestFile  time.Time
	AverageSize int64
}

// FormatSize returns a human-readable string representation of the total size.
func (s *DirStats) FormatSize() string {
	return FormatSize(s.TotalSize)
}

// FormatSize formats a byte count into a human-readable string (e.g., "1.5 MB").
func FormatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %ciB",
		float64(bytes)/float64(div), "KMGTPE"[exp])
}
