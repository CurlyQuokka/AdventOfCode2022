package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/hpsgame"
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

	game := hpsgame.NewGame(input)
	fmt.Printf("Score 1: %d\n", game.GetScore(true))
	fmt.Printf("Score 2: %d\n", game.GetScore(false))

	os.Exit(0)
}
