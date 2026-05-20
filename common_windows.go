//go:build windows

package files

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func getCreatedDate(info os.FileInfo) time.Time {
	created := info.ModTime()
	data, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if ok {
		created = time.Unix(0, data.CreationTime.Nanoseconds())
	}
	return created
}

func isExecutableFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	_, ok := fileDescriptions[ext]
	if ok {
		return fileDescriptions[ext].IsExecutable
	}
	return false
}
