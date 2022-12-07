package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	fullSize = 70_000_000
	requiredSize = 30_000_000
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	root := parseFileSystem(input)
	lessThan := findLessThan(root, 100_000)
	return strconv.Itoa(sumSize(lessThan))
}

func parseFileSystem(input string) folder {
	lines := strings.Split(input, "\n")	
	return parseFolder(&lines)
}

func findLessThan(root folder, n int) []folder {
	lessThan := []folder{}
	if root.Size() < n {
		lessThan = append(lessThan, root)
	}
	for _, child := range root.children {
		switch child := child.(type) {
		case *folder:
			lessThan = append(lessThan, findLessThan(*child, n)...)
		}
	}
	return lessThan
}

func sumSize(folders []folder) int {
	sum := 0
	for _, folder := range folders {
		sum += folder.Size()
	}
	return sum
}

type sizer interface {
	Size() int
	String() string
}

type file struct {
	name string
	size int
}

func newFile(name string, size int) *file {
	return &file{name: name, size: size}
}

func (f *file) Size() int {
	return f.size
}

func (f *file) String() string {
	return fmt.Sprintf("%s (%d)", f.name, f.size)
}

type folder struct {
	name string
	children []sizer
}

func newFolder(name string) *folder {
	return &folder{name: name, children: []sizer{}}
}

func parseFolder(lines *[]string) folder {
	firstLine := (*lines)[0]
	f := newFolder(firstLine[5:])
	*lines = (*lines)[2:]
	for ; 0 < len(*lines); *lines = (*lines)[1:] {
		line := (*lines)[0]
		switch {
		case line == "$ cd ..":
			return *f
		case strings.HasPrefix(line, "dir"):
			continue
		case strings.HasPrefix(line, "$ cd"):
			newFolder := parseFolder(lines) 
			f.Add(&newFolder)
		default:
			parts := strings.Split(line, " ")
			size, _ := strconv.Atoi(parts[0])
			f.Add(newFile(parts[1], size))
		}
		if len(*lines) == 0 {
			break
		}
	}
	return *f
}

func (f *folder) Add(child sizer) {
	f.children = append(f.children, child)
}

func (f *folder) Size() int {
	size := 0
	for _, child := range f.children {
		size += child.Size()
	}
	return size
}

func (f *folder) String() string {
	lines := make([]string, 1, len(f.children) + 1)
	lines[0] = fmt.Sprintf("%s (%d)", f.name, f.Size())
	for _, child := range f.children {
		childLines := strings.Split(child.String(), "\n")
		for j, line := range childLines {
			childLines[j] = fmt.Sprintf("  %s", line)
		}
		lines = append(lines, childLines...)
	}
	return strings.Join(lines, "\n")
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	root := parseFileSystem(input)
	freeSpace := fullSize - root.Size()
	sortedSizes := sortFolders(root)
	smallest := findSmallest(sortedSizes, requiredSize - freeSpace)
	return strconv.Itoa(smallest)
}

func sortFolders(root folder) []int {
	sortedSize := []int{}
	for _, child := range root.children {
		switch child := child.(type) {
		case *folder:
			sortedSize = append(sortedSize, sortFolders(*child)...)
		}
	}
	sortedSize = append(sortedSize, root.Size())
	sort.Ints(sortedSize)
	return sortedSize
}

func findSmallest(sizes []int, remaining int) int {
	for _, size := range sizes {
		if size > remaining {
			return size
		}
	}
	panic("No size found")
}
