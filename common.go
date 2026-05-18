package files

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

const maxDecompressedSize = 1 << 30 // 1 073 741 824 bytes
const bufferSize = 64 * 1024        // 64 KB buffer

// sanitizeFilePath prevents Zip Slip vulnerability (G305) and path traversal attacks (G703).
// It validates that the resulting path is within the base directory.
func sanitizeFilePath(base, subPath string) (string, error) {
	v := filepath.Join(base, subPath)
	if strings.HasPrefix(v, filepath.Clean(base)) {
		return v, nil
	}

	return "", fmt.Errorf("%w: %s", ErrPathIsTainted, subPath)
}

func copyIO(dst io.Writer, src io.Reader) error {
	var total int64
	buf := make([]byte, bufferSize)
	for {
		n, readErr := src.Read(buf)
		if n > 0 {
			total += int64(n)
			if total > maxDecompressedSize {
				return ErrFileIsTooBig
			}
			_, err := dst.Write(buf[:n])
			if err != nil {
				return err
			}
		}
		if readErr != nil {
			if errors.Is(readErr, io.EOF) {
				break
			}
			return readErr
		}
	}

	return nil
}
