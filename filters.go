package files

import (
	"regexp"
	"strings"
	"time"
)

// nolint: gochecknoglobals
var (
	// Directory filters.

	FilterDirByName = func(name string) FilterDir {
		return func(f *DirInfo) bool {
			return name == f.Name
		}
	}

	FilterDirByRegEx = func(regex regexp.Regexp) FilterDir {
		return func(f *DirInfo) bool {
			return regex.MatchString(f.Name)
		}
	}

	FilterDirBySizeGreater = func(size int64) FilterDir {
		return func(f *DirInfo) bool {
			return f.Size > size
		}
	}

	FilterDirBySizeLower = func(size int64) FilterDir {
		return func(f *DirInfo) bool {
			return f.Size < size
		}
	}

	FilterDirByNbFilesGreater = func(nb int64) FilterDir {
		return func(f *DirInfo) bool {
			return f.Nbfiles > nb
		}
	}

	FilterDirByNbFilesLower = func(nb int64) FilterDir {
		return func(f *DirInfo) bool {
			return f.Nbfiles < nb
		}
	}

	FilterDirByNbFilesEqual = func(nb int64) FilterDir {
		return func(f *DirInfo) bool {
			return f.Nbfiles == nb
		}
	}

	// File filters.

	FilterFileByName = func(name string) FilterFile {
		return func(f *FileInfo) bool {
			return name == f.Name
		}
	}

	FilterFileByRegEx = func(regex regexp.Regexp) FilterFile {
		return func(f *FileInfo) bool {
			return regex.MatchString(f.Name)
		}
	}

	FilterFileByExt = func(ext string) FilterFile {
		return func(f *FileInfo) bool {
			return strings.ToLower(ext) == f.GetExt()
		}
	}

	FilterFileBySizeGreater = func(size int64) FilterFile {
		return func(f *FileInfo) bool {
			return f.Size > size
		}
	}

	FilterFileBySizeLower = func(size int64) FilterFile {
		return func(f *FileInfo) bool {
			return f.Size < size
		}
	}

	FilterFileByCreatedAfter = func(datetime time.Time) FilterFile {
		return func(f *FileInfo) bool {
			return f.Created.UnixNano() > datetime.UnixNano()
		}
	}

	FilterFileByCreatedBefore = func(datetime time.Time) FilterFile {
		return func(f *FileInfo) bool {
			return f.Created.UnixNano() < datetime.UnixNano()
		}
	}

	FilterFileByUpdatedAfter = func(datetime time.Time) FilterFile {
		return func(f *FileInfo) bool {
			return f.Updated.UnixNano() > datetime.UnixNano()
		}
	}

	FilterFileByUpdatedBefore = func(datetime time.Time) FilterFile {
		return func(f *FileInfo) bool {
			return f.Updated.UnixNano() < datetime.UnixNano()
		}
	}
)
