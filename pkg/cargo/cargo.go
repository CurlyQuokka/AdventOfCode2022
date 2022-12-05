package cargo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

type Cargo struct {
	stacks      []*stack.Stack
	initial     []string
	operations  []string
	numOfStacks int
}

func NewCargo(input []string) (*Cargo, error) {
	breakingLine := getBreakingLine(input)
	if breakingLine < 0 {
		return nil, errors.New("Could not find breaking line")
	}

	numOfStacksSlice := strings.Split(strings.TrimSpace(input[breakingLine-1]), " ")
	numOfStacks, err := strconv.Atoi(numOfStacksSlice[len(numOfStacksSlice)-1])
	if err != nil {
		return nil, err
	}

	return &Cargo{
		initial:     input[:breakingLine],
		operations:  input[breakingLine+1:],
		numOfStacks: numOfStacks,
	}, nil
}

func (c *Cargo) InitializeStacks() {
	c.stacks = prepareStacks(c.numOfStacks)
	for i := len(c.initial) - 1; i >= 0; i-- {
		m := processStackLine(c.initial[i])
		for key, value := range m {
			c.stacks[key].Push(value)
		}
	}
}

func (c *Cargo) Peek() {
	for _, s := range c.stacks {
		fmt.Printf("%v", s.Peek())
	}
	fmt.Println()
}

func (c *Cargo) MoveCargo(moveMultipleAtOnce bool) error {
	for _, operation := range c.operations {
		operation = strings.ReplaceAll(operation, "move ", "")
		operation = strings.ReplaceAll(operation, "from ", "")
		operation = strings.ReplaceAll(operation, "to ", "")
		operations := strings.Split(operation, " ")
		operationsInts := []int{}
		for _, o := range operations {
			v, err := strconv.Atoi(o)
			if err != nil {
				return err
			}
			operationsInts = append(operationsInts, v)
		}
		if moveMultipleAtOnce {
			tmp := stack.New()
			for i := 0; i < operationsInts[0]; i++ {
				val := c.stacks[operationsInts[1]-1].Pop()
				tmp.Push(val)
			}
			for i := 0; i < operationsInts[0]; i++ {
				val := tmp.Pop()
				c.stacks[operationsInts[2]-1].Push(val)
			}
		} else {
			for i := 0; i < operationsInts[0]; i++ {
				val := c.stacks[operationsInts[1]-1].Pop()
				c.stacks[operationsInts[2]-1].Push(val)
			}
		}
	}
	return nil
}

func processStackLine(in string) map[int]string {
	m := make(map[int]string)
	for i, s := 0, 0; i < len(in); i, s = i+4, s+1 {
		val := strings.TrimSpace(in[i : i+3])
		val = strings.ReplaceAll(val, "[", "")
		val = strings.ReplaceAll(val, "]", "")
		if val != "" {
			m[s] = val
		}
	}
	return m
}

func prepareStacks(numOfStacks int) []*stack.Stack {
	s := []*stack.Stack{}
	for i := 0; i < numOfStacks; i++ {
		st := stack.New()
		s = append(s, st)
	}
	return s
}

func getBreakingLine(input []string) int {
	for i, line := range input {
		if line == "" {
			return i
		}
	}
	return -1
}
