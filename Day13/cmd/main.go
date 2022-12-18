package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/distress"
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

	d := distress.NewDistress(input)

	fmt.Println("Part 1:", d.GetPairSum())
	fmt.Println("Part 2:", d.GetDivPacketsMul())
}
