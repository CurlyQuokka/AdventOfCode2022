package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/elevation"
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

	e := elevation.NewElevation(input, false)
	fmt.Println("Shortest path - part 1: ", e.Bfs(elevation.EndingRune))

	e2 := elevation.NewElevation(input, true)
	fmt.Println("Shortest path - part 2: ", e2.Bfs(elevation.MinRune))
}
