package files_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alex-cos/files"
	"github.com/stretchr/testify/assert"
)

func TestCopyFile(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	err := os.MkdirAll(filepath.Join(tempdir, "testdata", "dummy"), 0755)
	assert.NoError(t, err)
	source := filepath.Join(".", "testdata", "dummy", "dummy1.txt")
	target := filepath.Join(tempdir, "testdata", "dummy", "dummy1.txt")

	err = files.CopyFile(source, target)
	assert.NoError(t, err)
}

func TestConcatFiles(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	err := os.MkdirAll(filepath.Join(tempdir, "testdata", "dummy"), 0755)
	assert.NoError(t, err)
	sources := []string{
		filepath.Join(".", "testdata", "dummy", "dummy1.txt"),
		filepath.Join(".", "testdata", "dummy", "dummy2.txt"),
	}
	target := filepath.Join(tempdir, "testdata", "dummy", "dummy3.txt")

	err = files.ConcatFiles(sources, target, 0755)
	assert.NoError(t, err)
}

func TestConcatDir(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	err := os.MkdirAll(filepath.Join(tempdir, "testdata", "dummy"), 0755)
	assert.NoError(t, err)
	source := filepath.Join(".", "testdata", "dummy")
	target := filepath.Join(tempdir, "testdata", "dummy", "dummy3.txt")

	err = files.ConcatDir(source, target, files.FilterFileByExt(files.TXT), 0755)
	assert.NoError(t, err)
}

func TestCopyDir(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	err := os.MkdirAll(filepath.Join(tempdir, "testdata", "dummy"), 0755)
	assert.NoError(t, err)
	source := filepath.Join(".", "testdata", "dummy")
	target := filepath.Join(tempdir, "testdata", "dummy3")

	err = files.CopyDir(source, target)
	assert.NoError(t, err)
}

func TestDeleteFile(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	testFile := filepath.Join(tempdir, "test.txt")

	err := os.WriteFile(testFile, []byte("content"), 0644)
	assert.NoError(t, err)

	_, err = os.Stat(testFile)
	assert.NoError(t, err)

	err = files.DeleteFile(testFile)
	assert.NoError(t, err)

	_, err = os.Stat(testFile)
	assert.True(t, os.IsNotExist(err))
}

func TestDeleteFileNotFound(t *testing.T) {
	t.Parallel()

	err := files.DeleteFile("/nonexistent/file.txt")
	assert.Error(t, err)
}

func TestDeleteFileIsDir(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()

	err := files.DeleteFile(tempdir)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "is a directory")
}

func TestDeleteDir(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	subdir := filepath.Join(tempdir, "subdir")
	err := os.MkdirAll(subdir, 0755)
	assert.NoError(t, err)

	testFile := filepath.Join(subdir, "test.txt")
	err = os.WriteFile(testFile, []byte("content"), 0644)
	assert.NoError(t, err)

	_, err = os.Stat(subdir)
	assert.NoError(t, err)

	err = files.DeleteDir(subdir)
	assert.NoError(t, err)

	_, err = os.Stat(subdir)
	assert.True(t, os.IsNotExist(err))
}

func TestDeleteDirNotFound(t *testing.T) {
	t.Parallel()

	err := files.DeleteDir("/nonexistent/dir")
	assert.Error(t, err)
}

func TestDeleteDirNotDir(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	testFile := filepath.Join(tempdir, "test.txt")

	err := os.WriteFile(testFile, []byte("content"), 0644)
	assert.NoError(t, err)

	err = files.DeleteDir(testFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not a directory")
}

func TestFileHash(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	testFile := filepath.Join(tempdir, "test.txt")
	content := []byte("hello world")

	err := os.WriteFile(testFile, content, 0644)
	assert.NoError(t, err)

	hash, err := files.FileHash(testFile, files.SHA256)
	assert.NoError(t, err)
	assert.Len(t, hash, 64)
}

func TestFileHashMD5(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	testFile := filepath.Join(tempdir, "test.txt")
	content := []byte("hello world")

	err := os.WriteFile(testFile, content, 0644)
	assert.NoError(t, err)

	hash, err := files.FileHash(testFile, files.MD5)
	assert.NoError(t, err)
	assert.Len(t, hash, 32)
}

func TestFileHashSHA1(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	testFile := filepath.Join(tempdir, "test.txt")
	content := []byte("hello world")

	err := os.WriteFile(testFile, content, 0644)
	assert.NoError(t, err)

	hash, err := files.FileHash(testFile, files.SHA1)
	assert.NoError(t, err)
	assert.Len(t, hash, 40)
}

func TestFileHashSHA512(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	testFile := filepath.Join(tempdir, "test.txt")
	content := []byte("hello world")

	err := os.WriteFile(testFile, content, 0644)
	assert.NoError(t, err)

	hash, err := files.FileHash(testFile, files.SHA512)
	assert.NoError(t, err)
	assert.Len(t, hash, 128)
}

func TestFileHashNotFound(t *testing.T) {
	t.Parallel()

	_, err := files.FileHash("/nonexistent/file.txt", files.SHA256)
	assert.Error(t, err)
}

func TestFileHashIsDir(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()

	_, err := files.FileHash(tempdir, files.SHA256)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "is a directory")
}

func TestFileHashUnsupportedAlgo(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	testFile := filepath.Join(tempdir, "test.txt")

	err := os.WriteFile(testFile, []byte("content"), 0644)
	assert.NoError(t, err)

	_, err = files.FileHash(testFile, "unsupported")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported algorithm")
}

func TestGetDirStats(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	subdir := filepath.Join(tempdir, "subdir")
	err := os.MkdirAll(subdir, 0755)
	assert.NoError(t, err)

	file1 := filepath.Join(subdir, "file1.txt")
	err = os.WriteFile(file1, []byte("hello"), 0644)
	assert.NoError(t, err)

	file2 := filepath.Join(subdir, "file2.txt")
	err = os.WriteFile(file2, []byte("world!"), 0644)
	assert.NoError(t, err)

	stats, err := files.GetDirStats(subdir)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), stats.TotalFiles)
	assert.Equal(t, int64(1), stats.TotalDirs)
	assert.Equal(t, int64(11), stats.TotalSize)
	assert.Equal(t, int64(5), stats.AverageSize)
}

func TestGetDirStatsNotFound(t *testing.T) {
	t.Parallel()

	_, err := files.GetDirStats("/nonexistent/dir")
	assert.Error(t, err)
}

func TestGetDirStatsNotDir(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	testFile := filepath.Join(tempdir, "test.txt")

	err := os.WriteFile(testFile, []byte("content"), 0644)
	assert.NoError(t, err)

	_, err = files.GetDirStats(testFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not a directory")
}

func TestGetDirStatsEmptyDir(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()

	stats, err := files.GetDirStats(tempdir)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), stats.TotalFiles)
	assert.Equal(t, int64(1), stats.TotalDirs)
	assert.Equal(t, int64(0), stats.TotalSize)
}
