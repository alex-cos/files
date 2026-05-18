package files_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/alex-cos/files"
	"github.com/stretchr/testify/assert"
)

func TestListDirs(t *testing.T) {
	t.Parallel()

	source := filepath.Join(".", "testdata")
	dirs, err := files.ListDirs(source, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, dirs)
	if !testing.Short() {
		for _, dir := range dirs {
			fmt.Fprintf(os.Stdout, "%s|%s|%s|%d\n", dir.Path, dir.Name, dir.FormatSize(), dir.Nbfiles)
		}
	}

	assert.Equal(t, []*files.DirInfo{
		{
			Path:    filepath.Join("testdata", "dummy", "directory1"),
			Name:    "directory1",
			NbDirs:  0,
			Nbfiles: 1,
			Size:    9016,
		},
		{
			Path:    filepath.Join("testdata", "dummy"),
			Name:    "dummy",
			NbDirs:  1,
			Nbfiles: 2,
			Size:    1060,
		},
		{
			Path:    `testdata`,
			Name:    "testdata",
			NbDirs:  1,
			Nbfiles: 0,
			Size:    0,
		}}, dirs)

	source = filepath.Join(".", "testdata", "xxxxxx")
	_, err = files.ListDirs(source, nil)
	assert.Error(t, err)

	source = filepath.Join(".", "testdata", "dummy", "dummy1.txt")
	_, err = files.ListDirs(source, nil)
	assert.Error(t, err)
}

func TestListFiles(t *testing.T) {
	t.Parallel()

	source := filepath.Join(".", "testdata", "dummy")
	items, err := files.ListFiles(source, nil)
	assert.NoError(t, err)
	assert.Len(t, items, 2)
	if !testing.Short() {
		for _, item := range items {
			fmt.Fprintf(os.Stdout, "%s|%s|%s|%s|%s|%s\n",
				item.Path,
				item.Name,
				item.GetExt(),
				item.FormatSize(),
				item.Created.UTC().Format(time.DateTime),
				item.Updated.UTC().Format(time.DateTime),
			)
		}
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}
	assert.Equal(t, filepath.Join("testdata", "dummy", "dummy1.txt"), items[0].Path)
	assert.Equal(t, `dummy1.txt`, items[0].Name)
	assert.Equal(t, `txt`, items[0].GetExt())
	assert.Equal(t, int64(536), items[0].Size)
	assert.Equal(t, filepath.Join("testdata", "dummy", "dummy2.txt"), items[1].Path)
	assert.Equal(t, `dummy2.txt`, items[1].Name)
	assert.Equal(t, `txt`, items[1].GetExt())
	assert.Equal(t, int64(524), items[1].Size)

	items, err = files.ListFiles(source, files.FilterFileByName("dummy1.txt"))
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.ListFiles(source, files.FilterFileByExt(files.TXT))
	assert.NoError(t, err)
	assert.Len(t, items, 2)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.ListFiles(source, files.FilterFileBySizeGreater(530))
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.ListFiles(source, files.FilterFileBySizeLower(530))
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	source = filepath.Join(".", "testdata", "xxxxxx")
	_, err = files.ListFiles(source, files.FilterFileAll())
	assert.Error(t, err)

	source = filepath.Join(".", "testdata", "dummy", "dummy1.txt")
	_, err = files.ListFiles(source, files.FilterFileAll())
	assert.Error(t, err)
}

func TestWalkFiles(t *testing.T) {
	t.Parallel()

	source := filepath.Join(".", "testdata")
	items, err := files.WalkFiles(source, nil)
	assert.NoError(t, err)
	assert.Len(t, items, 3)
	if !testing.Short() {
		for _, item := range items {
			fmt.Fprintf(os.Stdout, "%s|%s|%s|%s|%s|%s\n",
				item.Path,
				item.Name,
				item.GetExt(),
				item.FormatSize(),
				item.Created.UTC().Format(time.DateTime),
				item.Updated.UTC().Format(time.DateTime),
			)
		}
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}
	assert.Equal(t, filepath.Join("testdata", "dummy", "directory1", "logo.jpg"), items[0].Path)
	assert.Equal(t, `logo.jpg`, items[0].Name)
	assert.Equal(t, `jpg`, items[0].GetExt())
	assert.Equal(t, int64(9016), items[0].Size)
	assert.Equal(t, filepath.Join("testdata", "dummy", "dummy1.txt"), items[1].Path)
	assert.Equal(t, `dummy1.txt`, items[1].Name)
	assert.Equal(t, `txt`, items[1].GetExt())
	assert.Equal(t, int64(536), items[1].Size)
	assert.Equal(t, filepath.Join("testdata", "dummy", "dummy2.txt"), items[2].Path)
	assert.Equal(t, `dummy2.txt`, items[2].Name)
	assert.Equal(t, `txt`, items[2].GetExt())
	assert.Equal(t, int64(524), items[2].Size)

	items, err = files.WalkFiles(source, files.FilterFileByName("dummy1.txt"))
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.WalkFiles(source, files.FilterFileByExt(files.TXT))
	assert.NoError(t, err)
	assert.Len(t, items, 2)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.WalkFiles(source, files.FilterFileBySizeGreater(530))
	assert.NoError(t, err)
	assert.Len(t, items, 2)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.WalkFiles(source, files.FilterFileBySizeLower(530))
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	source = filepath.Join(".", "testdata", "xxxxxx")
	_, err = files.WalkFiles(source, files.FilterFileAll())
	assert.Error(t, err)

	source = filepath.Join(".", "testdata", "dummy", "dummy1.txt")
	_, err = files.WalkFiles(source, files.FilterFileAll())
	assert.Error(t, err)
}
