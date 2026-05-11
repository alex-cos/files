package files

import (
	"archive/tar"
	"errors"
	"io"
	"os"
	"path/filepath"
)

// TarList lists all files contained in a TAR archive.
func TarList(tarfile string) ([]string, error) {
	files := []string{}

	inFile, err := os.Open(tarfile)
	if err != nil {
		return files, err
	}
	defer inFile.Close()

	archive := tar.NewReader(inFile)
	for {
		hdr, err := archive.Next()
		if errors.Is(err, io.EOF) {
			break // End of archive
		}
		if err != nil {
			return files, err
		}
		if hdr == nil {
			continue
		}
		files = append(files, hdr.Name)
	}

	return files, nil
}

// UnTar extracts a TAR archive to the destination directory.
func UnTar(tarfile, dest string) error {
	inFile, err := os.Open(tarfile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	archive := tar.NewReader(inFile)
	for {
		hdr, err := archive.Next()
		if errors.Is(err, io.EOF) {
			break // End of archive
		}
		if err != nil {
			return err
		}
		if hdr == nil {
			continue
		}
		path, err := sanitizeFilePath(dest, hdr.Name)
		if err != nil {
			return err
		}
		if hdr.FileInfo().IsDir() {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, hdr.FileInfo().Mode())
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.CopyN(dstFile, archive, maxDecompressedSize)
		if err != nil {
			if errors.Is(err, io.EOF) {
				continue
			}
			return err
		}
	}

	return nil
}

// TarAll creates a TAR archive from a source folder.
func TarAll(baseFolder, tarfile string) error {
	outFile, err := os.Create(tarfile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := tar.NewWriter(outFile)

	err = addToTar(writer, baseFolder, "")
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	return nil
}

// addToTar recursively adds files from a directory to a TAR archive.
func addToTar(writer *tar.Writer, basePath, baseInTar string) error {
	files, err := os.ReadDir(basePath)
	if err != nil {
		return err
	}
	for _, file := range files {
		newBase := filepath.Join(basePath, file.Name())
		if !file.IsDir() {
			dat, err := os.ReadFile(newBase)
			if err != nil {
				return err
			}
			fileInfo, err := file.Info()
			if err != nil {
				return err
			}
			//nolint:exhaustruct
			hdr := &tar.Header{
				Name:    filepath.Join(baseInTar, file.Name()),
				Mode:    int64(fileInfo.Mode()),
				Size:    int64(len(dat)),
				ModTime: fileInfo.ModTime(),
			}
			err = writer.WriteHeader(hdr)
			if err != nil {
				return err
			}
			_, err = writer.Write(dat)
			if err != nil {
				return err
			}
		} else {
			err = addToTar(writer, newBase, filepath.Join(baseInTar, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
