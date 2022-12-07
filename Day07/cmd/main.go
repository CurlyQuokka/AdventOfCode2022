package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/filesystem"
	"github.com/CurlyQuokka/AdventOfCode2022/pkg/utils"
)

func main() {
	filePath := "input"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	fileSize := 100000
	if len(os.Args) > 2 {
		fSize, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fileSize = fSize
	}

	input, err := utils.LoadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fs := filesystem.NewFilesystem()
	err = fs.ProcessCommands(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Size of files smaller than %d: %d\n", fileSize, fs.GetSizeOfSmallerThan(fileSize))
	fmt.Printf("Size of directory to delete for update: %d\n", fs.GetSizeOfMinDir())
}
