package files

import (
	"fmt"
	"path/filepath"
	"strings"
)

const maxDecompressedSize = 1 << 20 // 1 Megabyte = 1024 * 1024 bytes

// Sanitize archive file pathing from "G305: Zip Slip vulnerability".
func sanitizeArchivePath(path, file string) (string, error) {
	v := filepath.Join(path, file)
	if strings.HasPrefix(v, filepath.Clean(path)) {
		return v, nil
	}

	return "", fmt.Errorf("content filepath is tainted: %s", file)
}
