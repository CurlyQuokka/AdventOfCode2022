package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/communication"
	"github.com/CurlyQuokka/AdventOfCode2022/pkg/utils"
)

func main() {
	filePath := "input"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	windowSize := 4
	if len(os.Args) > 2 {
		ws, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		windowSize = ws
	}

	input, err := utils.LoadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	chars := communication.ProcessSignal(input[0], windowSize)
	fmt.Println("Result:", chars)

	os.Exit(0)
}
