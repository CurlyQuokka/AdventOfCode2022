package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/utils"
	"github.com/CurlyQuokka/AdventOfCode2022/pkg/video"
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

	v := video.NewVideo()
	err = v.ProcessInput(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Signal strength: ", v.GetSignalStrength())
	v.ShowScreen()
}
