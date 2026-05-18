//go:build windows

package files

import (
	"os"
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
