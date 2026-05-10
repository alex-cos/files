package files

import (
	"fmt"
	"path/filepath"
	"strings"
)

const maxDecompressedSize = 1 << 20 // 1 Megabyte = 1024 * 1024 bytes

// Sanitize file pathing from Zip Slip vulnerability G305 and
// path traversal attacks G703.
func sanitizeFilePath(base, subPath string) (string, error) {
	v := filepath.Join(base, subPath)
	if strings.HasPrefix(v, filepath.Clean(base)) {
		return v, nil
	}

	return "", fmt.Errorf("filepath is tainted: %s", subPath)
}
