package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/rope"
	"github.com/CurlyQuokka/AdventOfCode2022/pkg/utils"
)

func main() {
	filePath := "input"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	numOfKnots := 2
	if len(os.Args) > 2 {
		nk, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		numOfKnots = nk
	}

	input, err := utils.LoadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	rope := rope.NewRope(numOfKnots)
	err = rope.ProcessInput(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of visited:", rope.CountVisited())
}
