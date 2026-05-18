//go:build !windows

package files

import (
	"os"
	"time"
)

func getCreatedDate(info os.FileInfo) time.Time {
	return info.ModTime()
}
