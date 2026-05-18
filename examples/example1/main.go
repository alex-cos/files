package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/alex-cos/files"
	"github.com/alexeyco/simpletable"
	"github.com/jftuga/ellipsis"
)

func main() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filter := files.FilterFileAll()

	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	if len(os.Args) > 2 {
		filter = files.FilterFileByExt(os.Args[2])
	}

	err := displayStats(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "")

	err = listFiles(dir, filter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func displayStats(dir string) error {
	stats, err := files.GetDirStats(dir)
	if err != nil {
		return err
	}

	table := simpletable.New()
	table.SetStyle(simpletable.StyleMarkdown)
	table.Body.Cells = [][]*simpletable.Cell{
		{
			{Align: simpletable.AlignLeft, Text: "PATH"},
			{Align: simpletable.AlignLeft, Text: filepath.Clean(dir)},
		},
		{
			{Align: simpletable.AlignLeft, Text: "TOTAL FILES"},
			{Align: simpletable.AlignLeft, Text: strconv.FormatInt(stats.TotalFiles, 10)},
		},
		{
			{Align: simpletable.AlignLeft, Text: "TOTAL DIRECTORIES"},
			{Align: simpletable.AlignLeft, Text: strconv.FormatInt(stats.TotalDirs, 10)},
		},
		{
			{Align: simpletable.AlignLeft, Text: "TOTAL SIZE"},
			{Align: simpletable.AlignLeft, Text: files.FormatSize(stats.TotalSize)},
		},
		{
			{Align: simpletable.AlignLeft, Text: "AVERAGE SIZE"},
			{Align: simpletable.AlignLeft, Text: files.FormatSize(stats.AverageSize)},
		},
		{
			{Align: simpletable.AlignLeft, Text: "OLDEST"},
			{Align: simpletable.AlignLeft, Text: stats.OldestFile.Format(time.DateTime)},
		},
		{
			{Align: simpletable.AlignLeft, Text: "NEWEST"},
			{Align: simpletable.AlignLeft, Text: stats.NewestFile.Format(time.DateTime)},
		},
	}
	fmt.Fprintln(os.Stdout, table.String())

	return nil
}

func listFiles(dir string, filter files.FilterFile) error {
	items, err := files.ListFiles(dir, filter)
	if err != nil {
		return err
	}

	table := simpletable.New()
	table.SetStyle(simpletable.StyleMarkdown)
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "PATH"},
			{Align: simpletable.AlignCenter, Text: "EXT"},
			{Align: simpletable.AlignCenter, Text: "SIZE"},
			{Align: simpletable.AlignCenter, Text: "CATEGORY"},
			{Align: simpletable.AlignCenter, Text: "CREATED"},
			{Align: simpletable.AlignCenter, Text: "UPDATED"},
		},
	}

	if len(items) == 0 {
		table.Body.Cells = append(table.Body.Cells,
			[]*simpletable.Cell{{Span: 6, Align: simpletable.AlignCenter, Text: "no files"}},
		)
	} else {
		for _, file := range items {
			line := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: ellipsis.Shorten(file.Path, 45)},
				{Align: simpletable.AlignLeft, Text: ellipsis.Shorten(file.GetExt(), 5)},
				{Align: simpletable.AlignRight, Text: file.FormatSize()},
				{Align: simpletable.AlignLeft, Text: files.GetFileDescCat(file.GetExt())},
				{Align: simpletable.AlignLeft, Text: file.Created.Format(time.DateTime)},
				{Align: simpletable.AlignLeft, Text: file.Updated.Format(time.DateTime)},
			}
			table.Body.Cells = append(table.Body.Cells, line)
		}
	}
	fmt.Fprintln(os.Stdout, table.String())

	return nil
}
