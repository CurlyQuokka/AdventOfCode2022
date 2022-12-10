package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/trees"
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

	trees, err := trees.NewTrees(input)
	// trees.PrintField()
	// fmt.Println()
	// trees.PrintVisiblility()
	// fmt.Println()
	trees.CheckVisibility()
	// trees.PrintVisiblility()
	// fmt.Println()
	fmt.Println("Visible trees:", trees.CountVisible())
	fmt.Println("Best scenic score:", trees.GetTopScenicScore())
}
