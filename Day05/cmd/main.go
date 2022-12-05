package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/cargo"
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

	c, err := cargo.NewCargo(input)
	if err != nil {
		log.Fatal(err)
	}

	c.InitializeStacks()
	err = c.MoveCargo(false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Part one: ")
	c.Peek()

	c.InitializeStacks()
	err = c.MoveCargo(true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Part two: ")
	c.Peek()

	os.Exit(0)
}
