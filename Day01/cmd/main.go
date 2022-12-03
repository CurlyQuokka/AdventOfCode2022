package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	values := []int{}
	currentValue := 0
	for scanner.Scan() {
		strVal := scanner.Text()
		if strVal != "" {
			v, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			currentValue += v
		} else {
			values = append(values, currentValue)
			currentValue = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if currentValue != 0 {
		values = append(values, currentValue)
	}

	sort.Ints(values)
	fmt.Println("Max cal:", values[len(values)-1])
	result := 0
	for i := len(values) - 1; i >= len(values)-topNum && i >= 0; i-- {
		result += values[i]
	}
	fmt.Printf("Top %d:\t %d\n", topNum, result)
}
