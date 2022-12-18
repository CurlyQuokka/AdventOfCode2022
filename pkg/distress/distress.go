package distress

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

const (
	divPacket1 = "[[2]]"
	divPacket2 = "[[6]]"
)

type Distress struct {
	pairs   []*pair
	packets []interface{}
}

func NewDistress(input []string) *Distress {
	d := &Distress{}
	for i := 0; i < len(input); i += 3 {
		d.pairs = append(d.pairs, newPair(input[i], input[i+1]))
	}
	d.Process()
	return d
}

func (d *Distress) Process() {
	for _, p := range d.pairs {
		p.process(d)
	}
	div1, _ := unmarshal(divPacket1)
	div2, _ := unmarshal(divPacket2)
	d.packets = append(d.packets, div1, div2)
	for _, p := range d.packets {
		fmt.Println(p)
	}
	sort.Slice(d.packets, d.sortCompare)
}

func (d *Distress) GetPairSum() int {
	sum := 0
	for i, p := range d.pairs {
		if p.valid {
			sum += (i + 1)
		}
	}
	return sum
}

func (d *Distress) GetDivPacketsMul() int {
	mul := 1
	for i, p := range d.packets {
		fmt.Println(p)
		if fmt.Sprint(p) == divPacket1 || fmt.Sprint(p) == divPacket2 {
			mul *= (i + 1)
		}
	}
	return mul
}

type pair struct {
	left, right string
	valid       bool
}

func newPair(left, right string) *pair {
	return &pair{
		left:  left,
		right: right,
		valid: false,
	}
}

func (p *pair) process(d *Distress) {
	left, errLeft := unmarshal(p.left)
	right, errRight := unmarshal(p.right)
	if errLeft != nil {
		log.Fatal(errLeft)
	}
	if errRight != nil {
		log.Fatal(errRight)
	}

	d.packets = append(d.packets, left, right)

	if compare(left, right) > 0 {
		p.valid = true
	}
}

func convertToList(data interface{}) ([]interface{}, bool) {
	res, isList := data.([]interface{})
	return res, isList
}

func unmarshal(data string) (interface{}, error) {
	var d interface{}
	err := json.Unmarshal([]byte(data), &d)
	return d, err
}

func (d *Distress) sortCompare(i, j int) bool {
	return compare(d.packets[i], d.packets[j]) > 0
}

func compare(left, right interface{}) int {
	leftCasted, leftOk := left.([]interface{})
	rightCasted, rightOk := right.([]interface{})

	switch {
	case !leftOk && !rightOk:
		return int(right.(float64) - left.(float64))
	case !leftOk:
		leftCasted = []interface{}{left}
	case !rightOk:
		rightCasted = []interface{}{right}
	}

	for i := 0; i < len(leftCasted) && i < len(rightCasted); i++ {
		if result := compare(leftCasted[i], rightCasted[i]); result != 0 {
			return result
		}
	}

	return len(rightCasted) - len(leftCasted)
}
