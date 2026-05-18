package files

import (
	"fmt"
	"os"
	"path/filepath"
)

// ListDirs lists all directories recursively within the given directory.
// It returns a slice of DirInfo for each directory found. If a filter is provided,
// only directories matching the filter are returned.
func ListDirs(directory string, filter FilterDir) ([]*DirInfo, error) {
	dirs := []*DirInfo{}
	dir := filepath.Clean(directory)

	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return dirs, fmt.Errorf("%w: %s", ErrDirNotFound, dir)
		}
		if os.IsPermission(err) {
			return dirs, fmt.Errorf("%w: %s", ErrPermissionDenied, dir)
		}
		return dirs, err
	}

	if !info.IsDir() {
		return dirs, fmt.Errorf("%w: %s", ErrDirIsFile, dir)
	}

	list, err := os.ReadDir(dir)
	if err != nil {
		return dirs, err
	}

	nbDirs := int64(0)
	nbFiles := int64(0)
	size := int64(0)
	for _, file := range list {
		if file.IsDir() {
			nbDirs++
			subdirs, err := ListDirs(filepath.Join(dir, file.Name()), filter)
			if err != nil {
				return dirs, err
			}
			dirs = append(dirs, subdirs...)
		} else {
			nbFiles++
			info, err := file.Info()
			if err != nil {
				continue
			}
			size += info.Size()
		}
	}

	dirInfo := &DirInfo{
		Path:    dir,
		Name:    filepath.Base(dir),
		NbDirs:  nbDirs,
		Nbfiles: nbFiles,
		Size:    size,
	}
	if filter != nil && !filter(dirInfo) {
		return dirs, nil
	}

	dirs = append(dirs, dirInfo)

	return dirs, nil
}

// ListFiles lists all files within the given directory.
// It returns a slice of FileInfo for each file found. If a filter is provided,
// only files matching the filter are returned.
func ListFiles(directory string, filter FilterFile) ([]*FileInfo, error) {
	files := []*FileInfo{}
	dir := filepath.Clean(directory)

	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return files, fmt.Errorf("%w: %s", ErrDirNotFound, dir)
		}
		if os.IsPermission(err) {
			return files, fmt.Errorf("%w: %s", ErrPermissionDenied, dir)
		}
		return files, err
	}

	if !info.IsDir() {
		return files, fmt.Errorf("%w: %s", ErrDirIsFile, dir)
	}

	list, err := os.ReadDir(dir)
	if err != nil {
		return files, err
	}

	for _, file := range list {
		if file.IsDir() {
			continue
		}

		info, err := file.Info()
		if err != nil {
			return files, err
		}

		fileInfo := &FileInfo{
			Path:    filepath.Join(dir, file.Name()),
			Name:    file.Name(),
			Ext:     filepath.Ext(file.Name()),
			Size:    info.Size(),
			Created: getCreatedDate(info),
			Updated: info.ModTime(),
		}
		if filter != nil && !filter(fileInfo) {
			continue
		}

		files = append(files, fileInfo)
	}

	return files, nil
}

// WalkFiles recursively walks through a directory and lists all files.
// It returns a slice of FileInfo for each file found. If a filter is provided,
// only files matching the filter are returned.
func WalkFiles(directory string, filter FilterFile) ([]*FileInfo, error) {
	files := []*FileInfo{}
	dir := filepath.Clean(directory)

	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return files, fmt.Errorf("%w: %s", ErrDirNotFound, dir)
		}
		if os.IsPermission(err) {
			return files, fmt.Errorf("%w: %s", ErrPermissionDenied, dir)
		}
		return files, err
	}

	if !info.IsDir() {
		return files, fmt.Errorf("%w: %s", ErrDirIsFile, dir)
	}

	err = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			fileInfo := &FileInfo{
				Path:    path,
				Name:    info.Name(),
				Ext:     filepath.Ext(path),
				Size:    info.Size(),
				Created: getCreatedDate(info),
				Updated: info.ModTime(),
			}

			if filter != nil && !filter(fileInfo) {
				return nil
			}

			files = append(files, fileInfo)

			return nil
		})

	if err != nil {
		return files, err
	}

	return files, nil
}
