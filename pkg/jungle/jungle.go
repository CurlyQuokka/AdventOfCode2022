package jungle

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	numberOfSeps       = 20
	monkeyLines        = 6
	operationSeparator = "="
	startingSeparator  = ":"
	valueSeparator     = ","
	spaceSeparator     = " "
	oldKeyWord         = "old"
)

type Jungle struct {
	monkeys          []*monkey
	commonDivider    int
	inspectionRelief int
}

func NewJungle(input []string, relief int) (*Jungle, error) {
	j := &Jungle{
		commonDivider:    1,
		inspectionRelief: relief,
	}
	for i := 0; i <= len(input)-monkeyLines; i += monkeyLines {
		m, err := newMonkey(input[i:i+monkeyLines], j)
		j.commonDivider *= m.divider
		if err != nil {
			return nil, err
		}
		j.monkeys = append(j.monkeys, m)
		if i+monkeyLines+1 < len(input) {
			i++
		}
	}
	return j, nil
}

func (j *Jungle) RunSimulation(steps int) {
	for i := 0; i < steps; i++ {
		for _, monkey := range j.monkeys {
			monkey.inspectItems()
		}
	}
}

func (j *Jungle) GetMonkeyBusiness() int {
	mb := []int{}
	for _, monkey := range j.monkeys {
		mb = append(mb, monkey.inspectionCounter)
	}
	sort.Ints(mb)
	return mb[len(mb)-1] * mb[len(mb)-2]
}

func (j *Jungle) getMonkey(index int) *monkey {
	return j.monkeys[index]
}

type monkey struct {
	jungle            *Jungle
	items             *list.List
	operation         string
	divider           int
	targets           []int
	inspectionCounter int
}

func newMonkey(data []string, j *Jungle) (*monkey, error) {
	m := &monkey{
		jungle: j,
		items:  list.New(),
	}
	items, err := getStartingItems(data[1])
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		m.addItem(item)
	}
	m.operation = getOperation(data[2])
	val, err := getLast(data[3])
	if err != nil {
		return nil, err
	}
	m.divider = val
	for i := 4; i < len(data); i++ {
		val, err = getLast(data[i])
		if err != nil {
			return nil, err
		}
		m.targets = append(m.targets, val)
	}

	return m, nil
}

func (m *monkey) addItem(value int) {
	m.items.PushBack(value)
}

func (m *monkey) throwItem(item *list.Element, target *monkey) error {
	v, ok := interface{}(item.Value).(int)
	if !ok {
		return fmt.Errorf("Cannot cast %v to int", item.Value)
	}
	target.addItem(v)
	m.items.Remove(item)
	return nil
}

func (m *monkey) inspectItems() error {
	var next *list.Element
	for item := m.items.Front(); item != nil; item = next {
		m.inspectionCounter++
		err := m.calculateNewWorryLvl(item)
		if err != nil {
			return err
		}
		next = item.Next()
		err = m.decideItemsFate(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *monkey) calculateNewWorryLvl(item *list.Element) error {
	opSplitted := strings.Split(m.operation, spaceSeparator)
	var a, b int
	var err error
	a, ok := interface{}(item.Value).(int)
	if !ok {
		return fmt.Errorf("Cannot cast %v to int", item.Value)
	}

	if opSplitted[2] == oldKeyWord {
		b = a
	} else {
		b, err = getLast(m.operation)
		if err != nil {
			return err
		}
	}
	var newValue int

	if opSplitted[1] == "*" {
		newValue = mutliply(a, b)
	} else {
		newValue = add(a, b)
	}

	item.Value = newValue % m.jungle.commonDivider / m.jungle.inspectionRelief
	return nil
}

func (m *monkey) decideItemsFate(item *list.Element) error {
	value, ok := interface{}(item.Value).(int)
	if !ok {
		return fmt.Errorf("Cannot cast %v to int", item.Value)
	}
	target := 0
	if value%m.divider != 0 {
		target = 1
	}
	targetMonkey := m.jungle.getMonkey(m.targets[target])
	m.throwItem(item, targetMonkey)
	return nil
}

func getOperation(data string) string {
	dataSplitted := strings.Split(data, operationSeparator)
	return strings.TrimSpace(dataSplitted[1])
}

func getStartingItems(data string) ([]int, error) {
	dataSplitted := strings.Split(data, startingSeparator)
	valuesStr := strings.ReplaceAll(dataSplitted[1], spaceSeparator, "")
	valuesStrSplitted := strings.Split(valuesStr, valueSeparator)
	values := []int{}
	for _, value := range valuesStrSplitted {
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		values = append(values, v)
	}
	return values, nil
}

func getLast(data string) (int, error) {
	dataSplitted := strings.Split(data, spaceSeparator)
	return strconv.Atoi(dataSplitted[len(dataSplitted)-1])
}

func mutliply(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}
