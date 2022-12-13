package rope

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/utils"
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
		if knot.currPos.X != 0 || knot.currPos.Y != 0 {
			vis[knot.currPos.Y+16][knot.currPos.X+12] = fmt.Sprintf("%d", k)
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

type knot struct {
	currPos utils.Position
	prevPos utils.Position
}

func (k *knot) update(command string) {
	k.prevPos = k.currPos
	switch command {
	case moveUp:
		k.currPos.Y--
	case moveDown:
		k.currPos.Y++
	case moveLeft:
		k.currPos.X--
	default:
		k.currPos.X++
	}

}

func (k *knot) updateWithKnot(h knot) {
	if calculateDistance(k.currPos, h.currPos) > updateLimit {
		k.prevPos = k.currPos
		if k.currPos.X == h.currPos.X || k.currPos.Y == h.currPos.Y {
			if k.currPos.X == h.currPos.X {
				if k.currPos.Y > h.currPos.Y {
					k.currPos.Y--
				} else {
					k.currPos.Y++
				}
			}
			if k.currPos.Y == h.currPos.Y {
				if k.currPos.X > h.currPos.X {
					k.currPos.X--
				} else {
					k.currPos.X++
				}
			}
		} else {
			if k.currPos.X < h.currPos.X {
				k.currPos.X++
			} else {
				k.currPos.X--
			}
			if k.currPos.Y < h.currPos.Y {
				k.currPos.Y++
			} else {
				k.currPos.Y--
			}
		}

	}
}

func (k *knot) getKey() string {
	return fmt.Sprintf("%d-%d", k.currPos.X, k.currPos.Y)
}

func calculateDistance(firstPos, secondPos utils.Position) float64 {
	return math.Sqrt(math.Pow(float64(secondPos.X-firstPos.X), 2) + math.Pow(float64(secondPos.Y-firstPos.Y), 2))
}
