package calories

import (
	"sort"
	"strconv"
)

type Calories struct {
	rank []int
}

func NewCalories(input []string) (*Calories, error) {
	c := &Calories{}
	currentValue := 0
	for i := 0; i < len(input); i++ {
		if input[i] != "" || i == len(input)-1 {
			v, err := strconv.Atoi(input[i])
			if err != nil {
				return nil, err
			}
			currentValue += v
		} else {
			c.rank = append(c.rank, currentValue)
			currentValue = 0
		}
	}
	sort.Ints(c.rank)
	return c, nil
}

func (c *Calories) GetTop(num int) int {
	sum := 0
	for i := len(c.rank) - 1; i >= len(c.rank)-num && i >= 0; i-- {
		sum += c.rank[i]
	}
	return sum
}
