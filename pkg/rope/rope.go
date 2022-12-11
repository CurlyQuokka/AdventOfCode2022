package rope

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	moveDown  = "D"
	moveUp    = "U"
	moveLeft  = "L"
	moveRight = "R"

	updateLimit = 1.42 // > sqrt(2)

	cmdSeparator = " "
)

type Rope struct {
	knots   []knot
	visited map[string]bool
}

func NewRope(length int) *Rope {
	knots := []knot{}
	for i := 0; i < length; i++ {
		knots = append(knots, knot{})
	}
	return &Rope{
		knots:   knots,
		visited: make(map[string]bool),
	}
}

func (r *Rope) ProcessInput(input []string) error {
	for _, command := range input {
		commandSplitted := strings.Split(command, cmdSeparator)
		steps, err := strconv.Atoi(commandSplitted[1])
		if err != nil {
			return err
		}

		for i := 0; i < steps; i++ {
			for k := 0; k < len(r.knots); k++ {
				if k == 0 {
					r.knots[k].update(commandSplitted[0])
				} else if k < len(r.knots)-1 {
					r.knots[k].updateWithKnot(r.knots[k-1])
				} else {
					r.knots[k].updateWithKnot(r.knots[k-1])
					key := r.knots[k].getKey()
					r.visited[key] = true
				}
			}
		}
	}
	return nil
}

func (r *Rope) PrintKnots() {
	vis := [30][30]string{}
	for row := range vis {
		for col := range vis[0] {
			vis[row][col] = "."
		}
	}
	vis[16][12] = "s"

	for k, knot := range r.knots {
		if knot.currPos.x != 0 || knot.currPos.y != 0 {
			vis[knot.currPos.y+16][knot.currPos.x+12] = fmt.Sprintf("%d", k)
		}
	}

	for _, row := range vis {
		fmt.Println(row)
	}
	fmt.Println("-----------------------")
}

func (r *Rope) CountVisited() int {
	return len(r.visited)
}

type position struct {
	x, y int
}

type knot struct {
	currPos position
	prevPos position
}

func (k *knot) update(command string) {
	k.prevPos = k.currPos
	switch command {
	case moveUp:
		k.currPos.y--
	case moveDown:
		k.currPos.y++
	case moveLeft:
		k.currPos.x--
	default:
		k.currPos.x++
	}

}

func (k *knot) updateWithKnot(h knot) {
	if calculateDistance(k.currPos, h.currPos) > updateLimit {
		k.prevPos = k.currPos
		if k.currPos.x == h.currPos.x || k.currPos.y == h.currPos.y {
			if k.currPos.x == h.currPos.x {
				if k.currPos.y > h.currPos.y {
					k.currPos.y--
				} else {
					k.currPos.y++
				}
			}
			if k.currPos.y == h.currPos.y {
				if k.currPos.x > h.currPos.x {
					k.currPos.x--
				} else {
					k.currPos.x++
				}
			}
		} else {
			if k.currPos.x < h.currPos.x {
				k.currPos.x++
			} else {
				k.currPos.x--
			}
			if k.currPos.y < h.currPos.y {
				k.currPos.y++
			} else {
				k.currPos.y--
			}
		}

	}
}

func (k *knot) getKey() string {
	return fmt.Sprintf("%d-%d", k.currPos.x, k.currPos.y)
}

func calculateDistance(firstPos, secondPos position) float64 {
	return math.Sqrt(math.Pow(float64(secondPos.x-firstPos.x), 2) + math.Pow(float64(secondPos.y-firstPos.y), 2))
}
