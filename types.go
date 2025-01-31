package files

import (
	"fmt"
	"strings"
	"time"
)

type DirInfo struct {
	Path    string
	Name    string
	Nbfiles int64
	Size    int64
}

type FilterDir func(f *DirInfo) bool

func (item *DirInfo) String() string {
	return fmt.Sprintf("%+v", *item)
}

func (item *DirInfo) FormatSize() string {
	return FormatSize(item.Size)
}

type FileInfo struct {
	Path    string
	Name    string
	Ext     string
	Size    int64
	Created time.Time
	Updated time.Time
}

func (item *FileInfo) String() string {
	return fmt.Sprintf("%+v", *item)
}

func (item *FileInfo) FormatSize() string {
	return FormatSize(item.Size)
}

func (item *FileInfo) GetExt() string {
	return strings.ToLower(strings.TrimLeft(item.Ext, "."))
}

type FilterFile func(f *FileInfo) bool

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
