package filesystem

import (
	"strconv"
	"strings"
)

const (
	rootName       = "/"
	upName         = ".."
	cmdSeparator   = " "
	cmdIndicator   = "$"
	dirIndicator   = "dir"
	filesystemSize = 70000000
	minimalSpace   = 30000000
)

type Filesystem struct {
	root       *directory
	currentDir *directory
}

func NewFilesystem() *Filesystem {
	rootDir := newDirectory(rootName, nil)
	return &Filesystem{
		root:       rootDir,
		currentDir: rootDir,
	}
}

func (fs *Filesystem) ProcessCommands(input []string) error {
	for i := 0; i < len(input); i++ {
		cmdSplit := strings.Split(input[i], cmdSeparator)
		if len(cmdSplit) > 2 {
			fs.commandCd(cmdSplit[2])
		} else {
			nextCommand := false
			arguments := []string{}
			for !nextCommand && i < len(input)-1 {
				i++
				if strings.Contains(input[i], cmdIndicator) {
					nextCommand = true
					i--
					break
				} else {
					arguments = append(arguments, input[i])
				}
			}
			if err := fs.commandLs(arguments); err != nil {
				return err
			}
		}
	}
	_ = fs.root.calculateSize()
	return nil
}

func (fs *Filesystem) GetSizeOfSmallerThan(limit int) int {
	return fs.root.getSizeOfSmallerThan(limit)
}

func (fs *Filesystem) GetSizeOfMinDir() int {
	limit := -1 * (filesystemSize - minimalSpace - fs.root.size)
	dirs := fs.root.getDirsByFilter(isGreaterOrEqual, limit)
	return findMinDirSize(dirs, limit)
}

func (fs *Filesystem) getDirsBiggerThan(limit int) []*directory {
	return fs.root.getDirsByFilter(isGreaterOrEqual, limit)
}

func (fs *Filesystem) commandCd(name string) {
	switch name {
	case upName:
		fs.currentDir = fs.currentDir.parent
	case rootName:
		fs.currentDir = fs.root
	default:
		for _, dir := range fs.currentDir.dirs {
			if dir.name == name {
				fs.currentDir = dir
				return
			}
		}
	}
}

func (fs *Filesystem) commandLs(data []string) error {
	for _, line := range data {
		splittedLine := strings.Split(line, cmdSeparator)
		if strings.Contains(splittedLine[0], dirIndicator) {
			fs.currentDir.addDir(splittedLine[1])
		} else {
			size, err := strconv.Atoi(splittedLine[0])
			if err != nil {
				return err
			}
			fs.currentDir.addFile(splittedLine[1], size)
		}
	}
	return nil
}

type file struct {
	name string
	size int
}

func newFile(name string, size int) *file {
	return &file{
		name: name,
		size: size,
	}
}

type directory struct {
	name   string
	parent *directory
	dirs   []*directory
	files  []*file
	size   int
}

func newDirectory(name string, parent *directory) *directory {
	return &directory{
		name:   name,
		parent: parent,
		dirs:   []*directory{},
		files:  []*file{},
		size:   0,
	}
}

func (d *directory) addDir(name string) {
	isExisting := false
	for _, dir := range d.dirs {
		if dir.name == name {
			isExisting = true
			break
		}
	}
	if !isExisting {
		d.dirs = append(d.dirs, newDirectory(name, d))
	}
}

func (d *directory) addFile(name string, size int) {
	isExisting := false
	for _, file := range d.files {
		if file.name == name {
			isExisting = true
			break
		}
	}
	if !isExisting {
		d.files = append(d.files, newFile(name, size))
	}
}

func (dir *directory) calculateSize() int {
	size := 0
	for _, f := range dir.files {
		size += f.size
	}
	for _, d := range dir.dirs {
		size += d.calculateSize()
	}
	dir.size = size
	return size
}

func (dir *directory) getDirsByFilter(filterFunc func(*directory, int) bool, limit int) []*directory {
	result := []*directory{}
	if filterFunc(dir, limit) {
		result = append(result, dir)
	}
	for _, d := range dir.dirs {
		result = append(result, d.getDirsByFilter(filterFunc, limit)...)
	}
	return result
}

func (dir *directory) getSizeOfSmallerThan(limit int) int {
	dirs := dir.getDirsByFilter(isSmallerOrEqual, limit)
	size := 0
	for _, d := range dirs {
		size += d.size
	}
	return size
}

func isSmallerOrEqual(d *directory, limit int) bool {
	return d.size <= limit
}

func isGreaterOrEqual(d *directory, limit int) bool {
	return d.size >= limit
}

func findMinDirSize(dirs []*directory, limit int) int {
	min := filesystemSize
	for _, dir := range dirs {
		if dir.size <= min && dir.size >= limit {
			min = dir.size
		}
	}
	return min
}
