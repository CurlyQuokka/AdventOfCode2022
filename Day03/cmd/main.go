package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/supplies"
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

	supplies := supplies.NewSupplies(input)
	fmt.Printf("Sum of priorities: %d\n", supplies.GetSumOfPriorites())

	fmt.Printf("Sum of group priorities: %d\n", supplies.GetSumOfGroupPriorities())

	os.Exit(0)
}
