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

	dirs, err := files.ListDirs("./testdata", nil)
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
			Nbfiles: 1,
			Size:    9016,
		},
		{
			Path:    filepath.Join("testdata", "dummy"),
			Name:    "dummy",
			Nbfiles: 2,
			Size:    1060,
		},
		{
			Path:    `testdata`,
			Name:    "testdata",
			Nbfiles: 0,
			Size:    0,
		}}, dirs)
}

func TestListFiles(t *testing.T) {
	t.Parallel()

	items, err := files.ListFiles("./testdata/dummy", nil)
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

	items, err = files.ListFiles("./testdata/dummy", files.FilterFileByName("dummy1.txt"))
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.ListFiles("./testdata/dummy", files.FilterFileByExt(files.TXT))
	assert.NoError(t, err)
	assert.Len(t, items, 2)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.ListFiles("./testdata/dummy", files.FilterFileBySizeGreater(530))
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.ListFiles("./testdata/dummy", files.FilterFileBySizeLower(530))
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}
}

func TestWalkFiles(t *testing.T) {
	t.Parallel()

	items, err := files.WalkFiles("./testdata", nil)
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

	items, err = files.WalkFiles("./testdata", files.FilterFileByName("dummy1.txt"))
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.WalkFiles("./testdata", files.FilterFileByExt(files.TXT))
	assert.NoError(t, err)
	assert.Len(t, items, 2)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.WalkFiles("./testdata", files.FilterFileBySizeGreater(530))
	assert.NoError(t, err)
	assert.Len(t, items, 2)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}

	items, err = files.WalkFiles("./testdata", files.FilterFileBySizeLower(530))
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	if !testing.Short() {
		fmt.Fprintf(os.Stdout, "files = %+v\n", items)
	}
}
