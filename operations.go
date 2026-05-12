package files

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// CopyFile copies a source file to a destination file or directory.
func CopyFile(src, dst string) error {
	var source = filepath.Clean(src)

	infosrc, err := os.Stat(source)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w: %s", ErrFileNotFound, source)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("%w: %s", ErrPermissionDenied, source)
		}
		return err
	}

	if infosrc.IsDir() {
		return fmt.Errorf("%w: %s", ErrFileIsDir, source)
	}

	destination, err := sanitizeFilePath(filepath.Dir(dst), filepath.Base(dst))
	if err != nil {
		return err
	}

	infodest, err := os.Stat(dst)
	switch {
	case err != nil:
		if !os.IsNotExist(err) {
			return err
		}
	case infodest.IsDir():
		destination = filepath.Join(destination, filepath.Base(source))
	default:
		return fmt.Errorf("%w: %s", ErrFileAlreadyExist, dst)
	}

	input, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	err = os.WriteFile(destination, input, infosrc.Mode())
	if err != nil {
		return err
	}

	return nil
}

// ConcatFiles concatenates all given files into one.
func ConcatFiles(sources []string, dst string, perm os.FileMode) error {
	var buffer []byte

	if len(sources) == 0 {
		return ErrEmptySource
	}

	for _, src := range sources {
		source := filepath.Clean(src)

		info, err := os.Stat(source)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("%w: %s", ErrFileNotFound, source)
			}
			if os.IsPermission(err) {
				return fmt.Errorf("%w: %s", ErrPermissionDenied, source)
			}
			return err
		}

		if info.IsDir() {
			return fmt.Errorf("%w: %s", ErrFileIsDir, src)
		}

		b, err := os.ReadFile(source)
		if err != nil {
			return err
		}
		buffer = append(buffer, b...)
	}

	destination, err := sanitizeFilePath(filepath.Dir(dst), filepath.Base(dst))
	if err != nil {
		return err
	}

	err = os.WriteFile(destination, buffer, perm)
	if err != nil {
		return err
	}

	return nil
}

// ConcatDir concatenates all files matching the given filter in a source directory.
func ConcatDir(src, dst string, filter FilterFile, perm os.FileMode) error {
	var source = filepath.Clean(src)
	var buffer []byte

	info, err := os.Stat(source)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w: %s", ErrDirNotFound, src)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("%w: %s", ErrPermissionDenied, src)
		}
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("%w: %s", ErrDirIsFile, src)
	}

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		cleanPath, err := sanitizeFilePath(source, path[len(source):])
		if err != nil {
			return err
		}

		fileInfo := &FileInfo{
			Path:    cleanPath,
			Name:    info.Name(),
			Ext:     filepath.Ext(cleanPath),
			Size:    info.Size(),
			Created: getCreatedDate(info),
			Updated: info.ModTime(),
		}

		if info.IsDir() || (filter != nil && !filter(fileInfo)) {
			return nil
		}

		b, e := os.ReadFile(cleanPath)
		if e != nil {
			return e
		}
		buffer = append(buffer, b...)

		return nil
	})

	if err != nil {
		return err
	}

	destination, err := sanitizeFilePath(filepath.Dir(dst), filepath.Base(dst))
	if err != nil {
		return err
	}

	err = os.WriteFile(destination, buffer, perm)
	if err != nil {
		return err
	}

	return nil
}

// CopyDir copies the entire source directory and sub-directories to a destination directory.
func CopyDir(src, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	info, err := os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w: %s", ErrDirNotFound, src)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("%w: %s", ErrPermissionDenied, src)
		}
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%w: %s", ErrDirIsFile, src)
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if err != nil {
		err = os.MkdirAll(dst, info.Mode())
		if err != nil {
			return err
		}
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			if entry.Type()&fs.ModeSymlink != 0 {
				continue
			}
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// DeleteFile removes a single file.
func DeleteFile(path string) error {
	source := filepath.Clean(path)

	info, err := os.Stat(source)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w: %s", ErrFileNotFound, source)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("%w: %s", ErrPermissionDenied, source)
		}
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("%w: %s", ErrFileIsDir, source)
	}

	return os.Remove(source)
}

// DeleteDir removes a directory and all its contents recursively.
func DeleteDir(path string) error {
	source := filepath.Clean(path)

	info, err := os.Stat(source)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w: %s", ErrDirNotFound, source)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("%w: %s", ErrPermissionDenied, source)
		}
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%w: %s", ErrDirIsFile, source)
	}

	return os.RemoveAll(source)
}

// MoveFile moves a file to a destination file or directory.
func MoveFile(src, dst string) error {
	source := filepath.Clean(src)

	info, err := os.Stat(source)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w: %s", ErrFileNotFound, source)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("%w: %s", ErrPermissionDenied, source)
		}
		return err
	}

	if info.IsDir() {
		return fmt.Errorf("%w: %s", ErrFileIsDir, source)
	}

	destination, err := sanitizeFilePath(filepath.Dir(dst), filepath.Base(dst))
	if err != nil {
		return err
	}

	infodest, err := os.Stat(dst)
	if err == nil && infodest.IsDir() {
		destination = filepath.Join(destination, filepath.Base(source))
	}

	err = os.Rename(source, destination)
	if err != nil {
		return err
	}

	return nil
}

// MoveDir moves a directory and all its contents to a destination directory.
func MoveDir(src, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	info, err := os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w: %s", ErrDirNotFound, src)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("%w: %s", ErrPermissionDenied, src)
		}
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("%w: %s", ErrDirIsFile, src)
	}

	dstInfo, err := os.Stat(dst)
	if err == nil && dstInfo.IsDir() {
		dst = filepath.Join(dst, filepath.Base(src))
	}

	err = os.Rename(src, dst)
	if err != nil {
		return err
	}

	return nil
}

// FileHash calculates the hash of a file using the specified algorithm (md5, sha1, sha256, sha512).
func FileHash(path, algo string) (string, error) {
	source := filepath.Clean(path)

	info, err := os.Stat(source)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("%w: %s", ErrFileNotFound, source)
		}
		if os.IsPermission(err) {
			return "", fmt.Errorf("%w: %s", ErrPermissionDenied, source)
		}
		return "", err
	}
	if info.IsDir() {
		return "", fmt.Errorf("%w: %s", ErrFileIsDir, source)
	}

	var h hash.Hash
	switch algo {
	case MD5:
		h = md5.New()
	case SHA1:
		h = sha1.New()
	case SHA256:
		h = sha256.New()
	case SHA512:
		h = sha512.New()
	default:
		return "", fmt.Errorf("unsupported algorithm: %s", algo)
	}

	file, err := os.Open(source)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}

	hashBytes := h.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

// GetDirStats calculates statistics for a directory (total files, size, oldest/newest file).
func GetDirStats(dir string) (*DirStats, error) {
	source := filepath.Clean(dir)

	info, err := os.Stat(source)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%w: %s", ErrDirNotFound, source)
		}
		if os.IsPermission(err) {
			return nil, fmt.Errorf("%w: %s", ErrPermissionDenied, source)
		}
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%w: %s", ErrDirIsFile, source)
	}

	var stats DirStats
	stats.OldestFile = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
	stats.NewestFile = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			stats.TotalDirs++
			return nil
		}

		stats.TotalFiles++
		stats.TotalSize += info.Size()

		modTime := info.ModTime()
		if modTime.Before(stats.OldestFile) {
			stats.OldestFile = modTime
		}
		if modTime.After(stats.NewestFile) {
			stats.NewestFile = modTime
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if stats.TotalFiles > 0 {
		stats.AverageSize = stats.TotalSize / stats.TotalFiles
	}

	return &stats, nil
}
