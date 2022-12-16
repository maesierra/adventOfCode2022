package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"maesierra.net/advent-of-code/2022/common"
)

type Day7 struct {
}

type Directory struct {
	name string
	parent *Directory
	directories map[string]*Directory
	files map[string]int
}

func (d Directory) AddFile(name string, size int) {
	d.files[name] = size	
}

func (d Directory) AddDirectory(name string) {
	d.directories[name] = &Directory{name, &d, map[string]*Directory{}, map[string]int{}}
}

func (d Directory) FilteredDirectoriesSize(maxSize int) int {
	size := 0
	if (d.Size() <= maxSize) {
		size += d.Size()
	}
	for _, d := range d.directories {
		size += d.FilteredDirectoriesSize(maxSize)
	}
	return size
}

func (d Directory) FindByMinSize(maxSize int) []Directory {
	res := make([]Directory, 0)
	if (d.Size() >= maxSize) {
		res = append(res, d)
	}
	for _, d := range d.directories {
		res = append(res, d.FindByMinSize(maxSize)...)
	}
	return res
}


func (d Directory) Size() int {
	size := 0
	for _, s := range d.files {
		size += s
	}
	for _, d := range d.directories {
		size += d.Size()
	}
	return size
}

func (d Directory) String() string {
	indent := ""
	for p := d.parent; p != nil ; p = p.parent {
		indent += " "
	}
	str := fmt.Sprintf("%s %s (dir, size=%d)\n", indent, d.name, d.Size())
	for _, d := range d.directories {
		str += fmt.Sprintf("%v", d)

	}
	for name, size := range d.files {
		str += fmt.Sprintf(" %s %v (file, size=%d)\n", indent, name, size)
	}
	return str
}

func (Day7) createFilesystem(inputFile string) Directory {
	root := Directory{
		"Filesystem",
		nil,
		map[string]*Directory{},
		map[string]int{},
	}
	root.AddDirectory("/")
	current := &root
	input := strings.Split(common.ReadFile(inputFile), "\n")
	for _, line := range input {
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == ".." {
					current = current.parent
				} else {
					current = current.directories[parts[2]]
				}
			}
		} else {
			if parts[0] == "dir" {
				current.AddDirectory(parts[1])
			} else {
				size, _ := strconv.Atoi(parts[0])
				current.AddFile(parts[1], size)
			}
		}

	}
	root = *root.directories["/"]
	root.parent = nil
	return root
}

func (d Day7) SolvePart1(inputFile string, data []string) string {		
	root := d.createFilesystem(inputFile)
	fmt.Printf("%v\n", root)
	total := root.FilteredDirectoriesSize(100000)
	return strconv.Itoa(total)

}


func (d Day7) SolvePart2(inputFile string, data []string) string {
	root := d.createFilesystem(inputFile)	
	diskSize := 70000000
	unUsed := diskSize - root.Size()
	candidates := root.FindByMinSize(30000000 - unUsed)
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i].Size() < candidates[j].Size()
	})
	return strconv.Itoa(candidates[0].Size())
}
