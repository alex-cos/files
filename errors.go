package files

import "errors"

var (
	// ErrFileNotFound is returned when a file does not exist.
	ErrFileNotFound = errors.New("file not found")

	// ErrDirNotFound is returned when a directory does not exist.
	ErrDirNotFound = errors.New("directory not found")

	// ErrFileIsDir is returned when a file path points to a directory.
	ErrFileIsDir = errors.New("path is a directory")

	// ErrDirIsFile is returned when a directory path points to a file.
	ErrDirIsFile = errors.New("path is a file")

	// ErrFileAlreadyExist is returned when a file already exist.
	ErrFileAlreadyExist = errors.New("file already exist")

	// ErrEmptySource is returned when no source files are provided.
	ErrEmptySource = errors.New("no source files provided")

	// ErrPermissionDenied is returned when file access is denied.
	ErrPermissionDenied = errors.New("permission denied")

	// ErrPathIsTainted is returned when a path is tainted.
	ErrPathIsTainted = errors.New("filepath is tainted")
)
