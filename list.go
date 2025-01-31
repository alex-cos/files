package files

import (
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
)

func ListDirectories(directory string, filter FilterDir) ([]*DirInfo, error) {
	dirs := []*DirInfo{}
	nbFiles := int64(0)
	size := int64(0)
	dir := filepath.Clean(directory)

	list, err := os.ReadDir(dir)
	if err != nil {
		return dirs, err
	}

	for _, file := range list {
		if file.IsDir() {
			subdirs, err := ListDirectories(filepath.Join(dir, file.Name()), filter)
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
		Nbfiles: nbFiles,
		Size:    size,
	}
	if filter != nil && !filter(dirInfo) {
		return dirs, nil
	}

	dirs = append(dirs, dirInfo)

	return dirs, nil
}

func ListFiles(directory string, filter FilterFile) ([]*FileInfo, error) {
	files := []*FileInfo{}

	list, err := os.ReadDir(directory)
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
			Path:    filepath.Join(directory, file.Name()),
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

func WalkFiles(directory string, filter FilterFile) ([]*FileInfo, error) {
	files := []*FileInfo{}

	err := filepath.Walk(
		directory,
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

func getCreatedDate(info os.FileInfo) time.Time {
	created := info.ModTime()
	if runtime.GOOS == "windows" {
		data, ok := info.Sys().(*syscall.Win32FileAttributeData)
		if ok {
			created = time.Unix(0, data.CreationTime.Nanoseconds())
		}
	} /* else if runtime.GOOS == "linux" {
		st := info.Sys().(*syscall.Stat_t)
		created = time.Unix(int64(st.Ctim.Sec), int64(st.Ctim.Nsec))
	}*/
	return created
}
