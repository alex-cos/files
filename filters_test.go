package files_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/alex-cos/files"
	"github.com/stretchr/testify/assert"
)

func TestFilterFileByName(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByName("test.txt")

	fileInfo := &files.FileInfo{
		Name: "test.txt",
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileByNameNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByName("test.txt")

	fileInfo := &files.FileInfo{
		Name: "other.txt",
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileByRegEx(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByRegEx(*regexp.MustCompile(`^test.*\.txt$`))

	fileInfo := &files.FileInfo{
		Name: "test_file.txt",
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileByRegExNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByRegEx(*regexp.MustCompile(`^test.*\.txt$`))

	fileInfo := &files.FileInfo{
		Name: "other.txt",
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileByExt(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByExt("go")

	const testFileName = "main.go"
	const testFileExt = ".go"

	fileInfo := &files.FileInfo{
		Name: testFileName,
		Ext:  testFileExt,
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileByExtCaseInsensitive(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByExt("GO")

	const testFileName = "main.go"
	const testFileExt = ".go"

	fileInfo := &files.FileInfo{
		Name: testFileName,
		Ext:  testFileExt,
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileByExtNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByExt("go")

	fileInfo := &files.FileInfo{
		Name: "main.txt",
		Ext:  ".txt",
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileByCategory(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByCategory(files.CategoryCode)

	fileInfo := &files.FileInfo{
		Name: "main.go",
		Ext:  ".go",
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileByCategoryNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByCategory(files.CategoryCode)

	fileInfo := &files.FileInfo{
		Name: "image.jpg",
		Ext:  ".jpg",
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileByCategoryUnknown(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByCategory(files.Unknown)

	fileInfo := &files.FileInfo{
		Name: "file.unknown",
		Ext:  ".unknown",
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileByMimeType(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByMimeType(files.MimeJSON)

	fileInfo := &files.FileInfo{
		Name: "test.json",
		Ext:  ".json",
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileByMimeTypeNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByMimeType(files.MimeJSON)

	fileInfo := &files.FileInfo{
		Name: "test.txt",
		Ext:  ".txt",
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileByMimeTypeUnknown(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByMimeType(files.MimePlain)

	fileInfo := &files.FileInfo{
		Name: "test.unknown",
		Ext:  ".unknown",
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileBySizeGreater(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileBySizeGreater(100)

	fileInfo := &files.FileInfo{
		Size: 200,
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileBySizeGreaterNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileBySizeGreater(100)

	fileInfo := &files.FileInfo{
		Size: 50,
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileBySizeLower(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileBySizeLower(100)

	fileInfo := &files.FileInfo{
		Size: 50,
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileBySizeLowerNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileBySizeLower(100)

	fileInfo := &files.FileInfo{
		Size: 200,
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileByCreatedAfter(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByCreatedAfter(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

	fileInfo := &files.FileInfo{
		Created: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileByCreatedAfterNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByCreatedAfter(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

	fileInfo := &files.FileInfo{
		Created: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileByCreatedBefore(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByCreatedBefore(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

	fileInfo := &files.FileInfo{
		Created: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileByCreatedBeforeNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByCreatedBefore(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

	fileInfo := &files.FileInfo{
		Created: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileByUpdatedAfter(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByUpdatedAfter(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

	fileInfo := &files.FileInfo{
		Updated: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileByUpdatedAfterNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByUpdatedAfter(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

	fileInfo := &files.FileInfo{
		Updated: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterFileByUpdatedBefore(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByUpdatedBefore(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

	fileInfo := &files.FileInfo{
		Updated: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	assert.True(t, filter(fileInfo))
}

func TestFilterFileByUpdatedBeforeNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterFileByUpdatedBefore(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

	fileInfo := &files.FileInfo{
		Updated: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	assert.False(t, filter(fileInfo))
}

func TestFilterDirByName(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirByName("mydir")

	dirInfo := &files.DirInfo{
		Name: "mydir",
	}

	assert.True(t, filter(dirInfo))
}

func TestFilterDirByNameNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirByName("mydir")

	dirInfo := &files.DirInfo{
		Name: "other",
	}

	assert.False(t, filter(dirInfo))
}

func TestFilterDirByRegEx(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirByRegEx(*regexp.MustCompile(`^dir.*`))

	dirInfo := &files.DirInfo{
		Name: "directory1",
	}

	assert.True(t, filter(dirInfo))
}

func TestFilterDirByRegExNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirByRegEx(*regexp.MustCompile(`^dir.*`))

	dirInfo := &files.DirInfo{
		Name: "other",
	}

	assert.False(t, filter(dirInfo))
}

func TestFilterDirBySizeGreater(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirBySizeGreater(1000)

	dirInfo := &files.DirInfo{
		Size: 2000,
	}

	assert.True(t, filter(dirInfo))
}

func TestFilterDirBySizeGreaterNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirBySizeGreater(1000)

	dirInfo := &files.DirInfo{
		Size: 500,
	}

	assert.False(t, filter(dirInfo))
}

func TestFilterDirBySizeLower(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirBySizeLower(1000)

	dirInfo := &files.DirInfo{
		Size: 500,
	}

	assert.True(t, filter(dirInfo))
}

func TestFilterDirBySizeLowerNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirBySizeLower(1000)

	dirInfo := &files.DirInfo{
		Size: 2000,
	}

	assert.False(t, filter(dirInfo))
}

func TestFilterDirByNbFilesGreater(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirByNbFilesGreater(5)

	dirInfo := &files.DirInfo{
		Nbfiles: 10,
	}

	assert.True(t, filter(dirInfo))
}

func TestFilterDirByNbFilesGreaterNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirByNbFilesGreater(5)

	dirInfo := &files.DirInfo{
		Nbfiles: 3,
	}

	assert.False(t, filter(dirInfo))
}

func TestFilterDirByNbFilesLower(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirByNbFilesLower(10)

	dirInfo := &files.DirInfo{
		Nbfiles: 5,
	}

	assert.True(t, filter(dirInfo))
}

func TestFilterDirByNbFilesLowerNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirByNbFilesLower(10)

	dirInfo := &files.DirInfo{
		Nbfiles: 15,
	}

	assert.False(t, filter(dirInfo))
}

func TestFilterDirByNbFilesEqual(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirByNbFilesEqual(5)

	dirInfo := &files.DirInfo{
		Nbfiles: 5,
	}

	assert.True(t, filter(dirInfo))
}

func TestFilterDirByNbFilesEqualNoMatch(t *testing.T) {
	t.Parallel()

	filter := files.FilterDirByNbFilesEqual(5)

	dirInfo := &files.DirInfo{
		Nbfiles: 10,
	}

	assert.False(t, filter(dirInfo))
}
