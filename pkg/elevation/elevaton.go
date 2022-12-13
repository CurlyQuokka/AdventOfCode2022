package elevation

import (
	"fmt"
	"sort"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/utils"
)

const (
	startingRune      = 'S'
	EndingRune        = 'E'
	MinRune           = 'a'
	maxRune           = 'z'
	maxElevtionChange = 1
)

type Elevation struct {
	elevationMap [][]rune
	startNode    *node
	endNode      *node
	nodes        map[string]*node
}

func NewElevation(input []string, reverted bool) *Elevation {
	e := &Elevation{
		nodes: make(map[string]*node),
	}
	for _, line := range input {
		row := []rune{}
		for _, item := range line {
			row = append(row, item)
		}
		e.elevationMap = append(e.elevationMap, row)
	}
	e.initializeNodes()
	e.updateConnections(reverted)
	if reverted {
		e.startNode.text = MinRune
		e.startNode = e.endNode
	}
	return e
}

func (e *Elevation) FindPathUsingBFS(endingSymbol rune) int {
	q := &queue{}
	pq := &pathQueue{}
	e.startNode.visited = true
	q.enqueue(e.startNode)
	var v *node
	path := []*node{}
	pq.enqueue(path)
	for q.len() > 0 {
		v = q.dequeue()
		p := pq.dequeue()
		p = append(p, v)
		if v.text == endingSymbol {
			return len(p) - 1
		}
		for _, con := range v.connected {
			if !con.visited {
				con.visited = true
				q.enqueue(con)
				pq.enqueue(p)
			}
		}
	}
	return -1
}

func (e *Elevation) initializeNodes() {
	for row := 0; row < len(e.elevationMap); row++ {
		for col := 0; col < len(e.elevationMap[0]); col++ {
			node := newNode(row, col, e.elevationMap[row][col])
			e.nodes[node.getKey()] = node
			if node.text == startingRune {
				e.startNode = node
			}
			if node.text == EndingRune {
				e.endNode = node
			}
		}
	}
}

func (e *Elevation) updateConnections(reverted bool) {
	for _, node := range e.nodes {
		node.updateConnections(e, reverted)
	}
}

func (e *Elevation) GetShortestPath(paths [][]*node) int {
	if len(paths) < 0 {
		return -1
	} else if len(paths) == 1 {
		return len(paths[0])
	}

	lengths := []int{}
	for _, p := range paths {
		lengths = append(lengths, len(p)-1)
	}
	sort.Ints(lengths)
	return lengths[0]
}

type node struct {
	pos       utils.Position
	text      rune
	value     int
	connected []*node
	visited   bool
}

func newNode(row, col int, text rune) *node {
	var value int
	switch text {
	case startingRune:
		value = MinRune
	case EndingRune:
		value = maxRune
	default:
		value = int(text)
	}
	return &node{
		pos:     utils.Position{X: col, Y: row},
		text:    text,
		value:   value,
		visited: false,
	}
}

func (n *node) getKey() string {
	return fmt.Sprintf("%d-%d", n.pos.Y, n.pos.X)
}

func (n *node) updateConnections(elevation *Elevation, reverted bool) {
	possible := n.createConnectionKeys()
	for _, key := range possible {
		dest, ok := elevation.nodes[key]
		if !ok {
			continue
		}
		if !reverted {
			if fromStart(n, dest) {
				n.connected = append(n.connected, dest)
			}
		} else {
			if fromEnd(n, dest) {
				n.connected = append(n.connected, dest)
			}
		}
	}
}

func (n *node) createConnectionKeys() []string {
	keys := []string{}
	keys = append(keys, fmt.Sprintf("%d-%d", n.pos.Y+1, n.pos.X))
	keys = append(keys, fmt.Sprintf("%d-%d", n.pos.Y, n.pos.X+1))
	keys = append(keys, fmt.Sprintf("%d-%d", n.pos.Y-1, n.pos.X))
	keys = append(keys, fmt.Sprintf("%d-%d", n.pos.Y, n.pos.X-1))
	return keys
}

type queue struct {
	items []*node
}

func (q *queue) enqueue(n *node) {
	q.items = append(q.items, n)
}

func (q *queue) dequeue() *node {
	if len(q.items) <= 0 {
		return nil
	}
	n := q.items[0]
	q.items = q.items[1:]
	return n
}

func (q *queue) len() int {
	return len(q.items)
}

type pathQueue struct {
	items [][]*node
}

func (pq *pathQueue) enqueue(n []*node) {
	pq.items = append(pq.items, n)
}

func (pq *pathQueue) dequeue() []*node {
	if len(pq.items) <= 0 {
		return nil
	}
	n := pq.items[0]
	pq.items = pq.items[1:]
	return n
}

func fromStart(src *node, dst *node) bool {
	return dst.value-src.value <= maxElevtionChange
}

func fromEnd(src *node, dst *node) bool {
	return src.value-dst.value <= maxElevtionChange
}
