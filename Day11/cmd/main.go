package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/jungle"
	"github.com/CurlyQuokka/AdventOfCode2022/pkg/utils"
)

const ()

func main() {
	filePath := "input"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	rounds := 20
	if len(os.Args) > 2 {
		nr, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		rounds = nr
	}

	relief := 3
	if len(os.Args) > 3 {
		rf, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		relief = rf
	}

	input, err := utils.LoadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	j, err := jungle.NewJungle(input, relief)
	if err != nil {
		log.Fatal(err)
	}

	j.RunSimulation(rounds)
	fmt.Println("Monkey bussiness level:", j.GetMonkeyBusiness())
}
