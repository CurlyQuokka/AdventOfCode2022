package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/sandcave"
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

	sc := sandcave.NewSandcave(input, false)
	sc.Simulate()
	fmt.Println("Part 1:", sc.CountSand())

	sc = sandcave.NewSandcave(input, true)
	sc.Simulate()
	fmt.Println("Part 2:", sc.CountSand())
}
