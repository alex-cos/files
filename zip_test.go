package files_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alex-cos/files"
	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	if !testing.Short() {
		fmt.Printf("tempdir = %+v\n", tempdir)
	}

	err := os.MkdirAll(filepath.Join(tempdir, "testdata"), 0755)
	assert.NoError(t, err)
	targetFile := filepath.Join(tempdir, "testdata", "dummy.zip")

	err = files.ZipAll("./testdata/dummy", targetFile)
	assert.NoError(t, err)

	items, err := files.ZipList(targetFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, items)
	if !testing.Short() {
		fmt.Printf("dirs = %+v\n", items)
	}

	err = files.UnZip(targetFile, filepath.Join(tempdir, "unzip"))
	assert.NoError(t, err)
}
