package files

import (
	"fmt"
	"path/filepath"
	"strings"
)

const maxDecompressedSize = 1 << 20 // 1 Megabyte = 1024 * 1024 bytes

// sanitizeFilePath prevents Zip Slip vulnerability (G305) and path traversal attacks (G703).
// It validates that the resulting path is within the base directory.
func sanitizeFilePath(base, subPath string) (string, error) {
	v := filepath.Join(base, subPath)
	if strings.HasPrefix(v, filepath.Clean(base)) {
		return v, nil
	}

	return "", fmt.Errorf("%w: %s", ErrPathIsTainted, subPath)
}
