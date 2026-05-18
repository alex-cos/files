package files_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alex-cos/files"
	"github.com/stretchr/testify/assert"
)

func TestTar(t *testing.T) {
	t.Parallel()

	tempdir := t.TempDir()
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "tempdir = %+v\n", tempdir)
	}

	err := os.MkdirAll(filepath.Join(tempdir, "testdata"), 0755)
	assert.NoError(t, err)
	targetFile := filepath.Join(tempdir, "testdata", "dummy.tar")

	err = files.TarAll(filepath.Join(".", "testdata", "dummy"), targetFile)
	assert.NoError(t, err)

	items, err := files.TarList(targetFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, items)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "dirs = %+v\n", items)
	}

	err = files.UnTar(targetFile, filepath.Join(tempdir, "untar"))
	assert.NoError(t, err)
}
