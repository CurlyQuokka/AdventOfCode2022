package cleaning

import (
	"strconv"
	"strings"
)

const (
	assignmentDelim = "-"
	pairDelim       = ","
)

type cleaning struct {
	pairs []pair
}

type assignment struct {
	startSection int
	endSection   int
}

type assignmentOverFunc func(assignment, assignment) bool

type pair struct {
	first  assignment
	second assignment
}

func NewCleaning(input []string) (*cleaning, error) {
	c := &cleaning{
		pairs: []pair{},
	}
	for _, line := range input {
		p, err := newPair(line)
		if err != nil {
			return nil, err
		}
		c.pairs = append(c.pairs, *p)
	}

	return c, nil
}

func (c *cleaning) GetOverlappedNumber(assignmentOverlapFunc assignmentOverFunc) uint {
	var counter uint
	counter = 0
	for _, pair := range c.pairs {
		if pair.isOverlapped(assignmentOverlapFunc) {
			counter++
		}
	}
	return counter
}

func newAssignment(in string) (*assignment, error) {
	sections := strings.Split(in, assignmentDelim)
	start, err := strconv.Atoi(sections[0])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(sections[1])
	if err != nil {
		return nil, err
	}
	return &assignment{
		startSection: start,
		endSection:   end,
	}, nil
}

func newPair(in string) (*pair, error) {
	assignments := strings.Split(in, pairDelim)
	first, err := newAssignment(assignments[0])
	if err != nil {
		return nil, err
	}
	second, err := newAssignment(assignments[1])
	if err != nil {
		return nil, err
	}
	return &pair{
		first:  *first,
		second: *second,
	}, nil
}

func (p *pair) isOverlapped(assignmentOverlapFunc assignmentOverFunc) bool {
	if assignmentOverlapFunc(p.first, p.second) || assignmentOverlapFunc(p.second, p.first) {
		return true
	}
	return false
}

func AreAssignemntsFullyOverlapped(a, b assignment) bool {
	if a.startSection >= b.startSection && a.endSection <= b.endSection {
		return true
	}
	return false
}

func AreAssignemntsOverlapped(a, b assignment) bool {
	if a.startSection >= b.startSection && a.startSection <= b.endSection || a.endSection >= b.startSection && a.endSection <= b.endSection {
		return true
	}
	return false
}
