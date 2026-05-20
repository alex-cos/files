//go:build !windows

package files

import (
	"os"
	"time"
)

func getCreatedDate(info os.FileInfo) time.Time {
	return info.ModTime()
}

func isExecutableFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode()&0111 != 0 // bit exécutable user/group/other
}
