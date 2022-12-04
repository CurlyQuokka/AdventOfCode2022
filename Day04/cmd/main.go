package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/cleaning"
	"github.com/CurlyQuokka/AdventOfCode2022/pkg/utils"
)

func main() {
	filePath := "input"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	input, err := utils.LoadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	c, err := cleaning.NewCleaning(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of full overlaps: %d\n", c.GetOverlappedNumber(cleaning.AreAssignmentsFullyOverlapped))
	fmt.Printf("Number of overlaps: %d\n", c.GetOverlappedNumber(cleaning.AreAssignmentsOverlapped))

	os.Exit(0)
}
