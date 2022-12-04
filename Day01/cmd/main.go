package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/calories"
	"github.com/CurlyQuokka/AdventOfCode2022/pkg/utils"
)

func main() {
	filePath := "input"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	topNum := 3
	if len(os.Args) > 2 {
		tn, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		topNum = tn
	}

	input, err := utils.LoadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	c, err := calories.NewCalories(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Max calories:\t", c.GetTop(1))
	fmt.Printf("Top %d:\t\t%d\n", topNum, c.GetTop(topNum))
}
