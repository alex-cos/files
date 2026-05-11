package files

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
)

// ZipList lists all files contained in a ZIP archive.
func ZipList(zipfile string) ([]string, error) {
	files := []string{}

	archive, err := zip.OpenReader(zipfile)
	if err != nil {
		return files, err
	}
	defer archive.Close()

	for _, f := range archive.File {
		files = append(files, f.Name)
	}

	return files, nil
}

// UnZip extracts a ZIP archive to the destination directory.
func UnZip(zipfile, dest string) error {
	archive, err := zip.OpenReader(zipfile)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, file := range archive.File {
		path, err := sanitizeFilePath(dest, file.Name)
		if err != nil {
			return err
		}
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer dstFile.Close()

		fileInArchive, err := file.Open()
		if err != nil {
			return err
		}
		defer fileInArchive.Close()

		_, err = io.CopyN(dstFile, fileInArchive, maxDecompressedSize)
		if err != nil {
			if errors.Is(err, io.EOF) {
				continue
			}
			return err
		}
	}

	return nil
}

// ZipAll creates a ZIP archive from a source folder.
func ZipAll(baseFolder, zipfile string) error {
	outFile, err := os.Create(zipfile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := zip.NewWriter(outFile)

	err = addToZip(writer, baseFolder, "")
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	return nil
}

// addToZip recursively adds files from a directory to a ZIP archive.
func addToZip(writer *zip.Writer, basePath, baseInZip string) error {
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
			f, err := writer.Create(filepath.Join(baseInZip, file.Name()))
			if err != nil {
				return err
			}
			_, err = f.Write(dat)
			if err != nil {
				return err
			}
		} else {
			_, err := writer.Create(filepath.Join(baseInZip, file.Name()) + "/")
			if err != nil {
				return err
			}
			err = addToZip(writer, newBase, filepath.Join(baseInZip, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
