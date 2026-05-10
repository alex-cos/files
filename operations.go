package files

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// CopyFile - copy a source file to a destination file or directory.
func CopyFile(src, dst string) error {
	var source = filepath.Clean(src)

	infosrc, err := os.Stat(source)
	if err != nil {
		return err
	}

	if infosrc.IsDir() {
		return fmt.Errorf("source file '%s' is a directory", source)
	}

	destination, err := sanitizeFilePath(filepath.Dir(dst), filepath.Base(dst))
	if err != nil {
		return err
	}

	infodest, err := os.Stat(dst)
	if err == nil && infodest.IsDir() {
		destination = filepath.Join(destination, filepath.Base(source))
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

// ConcatFiles - concatenate all given files into one.
func ConcatFiles(sources []string, dst string, perm os.FileMode) error {
	var buffer []byte

	for _, src := range sources {
		source := filepath.Clean(src)

		info, err := os.Stat(source)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return fmt.Errorf("source file '%s' is a directory", src)
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

// ConcatDir - concatenate all files of the same type located in a source directory.
func ConcatDir(src, dst string, filter FilterFile, perm os.FileMode) error {
	var source = filepath.Clean(src)
	var buffer []byte

	info, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("source file '%s' is not a directory", src)
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

// CopyDir - copy the entires source directory and sub-directories to a destination directory.
func CopyDir(src, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("source file '%s' is not a directory", src)
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
			// Skip symlinks.
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
