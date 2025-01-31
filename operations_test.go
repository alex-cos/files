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
