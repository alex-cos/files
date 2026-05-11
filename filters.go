package files

import (
	"regexp"
	"strings"
	"time"
)

// nolint: gochecknoglobals
var (
	// Directory filters.

	// FilterDirByName returns a filter that matches directories by exact name.
	FilterDirByName = func(name string) FilterDir {
		return func(f *DirInfo) bool {
			return name == f.Name
		}
	}

	// FilterDirByRegEx returns a filter that matches directories by regex pattern.
	FilterDirByRegEx = func(regex regexp.Regexp) FilterDir {
		return func(f *DirInfo) bool {
			return regex.MatchString(f.Name)
		}
	}

	// FilterDirBySizeGreater returns a filter that matches directories larger than the given size.
	FilterDirBySizeGreater = func(size int64) FilterDir {
		return func(f *DirInfo) bool {
			return f.Size > size
		}
	}

	// FilterDirBySizeLower returns a filter that matches directories smaller than the given size.
	FilterDirBySizeLower = func(size int64) FilterDir {
		return func(f *DirInfo) bool {
			return f.Size < size
		}
	}

	// FilterDirByNbFilesGreater returns a filter that matches directories with more files than the given count.
	FilterDirByNbFilesGreater = func(nb int64) FilterDir {
		return func(f *DirInfo) bool {
			return f.Nbfiles > nb
		}
	}

	// FilterDirByNbFilesLower returns a filter that matches directories with fewer files than the given count.
	FilterDirByNbFilesLower = func(nb int64) FilterDir {
		return func(f *DirInfo) bool {
			return f.Nbfiles < nb
		}
	}

	// FilterDirByNbFilesEqual returns a filter that matches directories with exactly the given number of files.
	FilterDirByNbFilesEqual = func(nb int64) FilterDir {
		return func(f *DirInfo) bool {
			return f.Nbfiles == nb
		}
	}

	// File filters.

	// FilterFileByName returns a filter that matches files by exact name.
	FilterFileByName = func(name string) FilterFile {
		return func(f *FileInfo) bool {
			return name == f.Name
		}
	}

	// FilterFileByRegEx returns a filter that matches files by regex pattern.
	FilterFileByRegEx = func(regex regexp.Regexp) FilterFile {
		return func(f *FileInfo) bool {
			return regex.MatchString(f.Name)
		}
	}

	// FilterFileByExt returns a filter that matches files by extension.
	FilterFileByExt = func(ext string) FilterFile {
		return func(f *FileInfo) bool {
			return strings.ToLower(ext) == f.GetExt()
		}
	}

	// FilterFileBySizeGreater returns a filter that matches files larger than the given size.
	FilterFileBySizeGreater = func(size int64) FilterFile {
		return func(f *FileInfo) bool {
			return f.Size > size
		}
	}

	// FilterFileBySizeLower returns a filter that matches files smaller than the given size.
	FilterFileBySizeLower = func(size int64) FilterFile {
		return func(f *FileInfo) bool {
			return f.Size < size
		}
	}

	// FilterFileByCreatedAfter returns a filter that matches files created after the given time.
	FilterFileByCreatedAfter = func(datetime time.Time) FilterFile {
		return func(f *FileInfo) bool {
			return f.Created.UnixNano() > datetime.UnixNano()
		}
	}

	// FilterFileByCreatedBefore returns a filter that matches files created before the given time.
	FilterFileByCreatedBefore = func(datetime time.Time) FilterFile {
		return func(f *FileInfo) bool {
			return f.Created.UnixNano() < datetime.UnixNano()
		}
	}

	// FilterFileByUpdatedAfter returns a filter that matches files updated after the given time.
	FilterFileByUpdatedAfter = func(datetime time.Time) FilterFile {
		return func(f *FileInfo) bool {
			return f.Updated.UnixNano() > datetime.UnixNano()
		}
	}

	// FilterFileByUpdatedBefore returns a filter that matches files updated before the given time.
	FilterFileByUpdatedBefore = func(datetime time.Time) FilterFile {
		return func(f *FileInfo) bool {
			return f.Updated.UnixNano() < datetime.UnixNano()
		}
	}
)
