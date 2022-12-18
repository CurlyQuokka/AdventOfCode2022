package sandcave

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2022/pkg/utils"
)

const (
	voidVal = 0
	sandVal = 1
	rockVal = 2

	dataSeparator     = " -> "
	positionSeparator = ","

	stateInMove  = 0
	stateStalled = 1
	stateOut     = 2

	maxSize = 1000
	startX  = 500
	startY  = 0
)

type Sandcave struct {
	scan        [maxSize][maxSize]int
	s           *sand
	notInfinite bool
}

func (sc *Sandcave) Simulate() {
	for {
		sc.s.move(sc)
		if sc.s.state == stateOut {
			break
		}
		if sc.s.state == stateStalled {
			sc.scan[sc.s.Y][sc.s.X] = sandVal
			if sc.s.Y == startY && sc.s.X == startX {
				break
			}
			sc.s.reset()
		}
	}
}

func NewSandcave(input []string, notInfinite bool) *Sandcave {
	sc := &Sandcave{
		s:           newSand(),
		notInfinite: notInfinite,
	}

	for _, data := range input {
		dataSplitted := strings.Split(data, dataSeparator)
		for i := 0; i < len(dataSplitted)-1; i++ {
			sc.drawLine(dataSplitted[i], dataSplitted[i+1])
		}
	}

	if sc.notInfinite {
		lowest := -1
		for i, row := range sc.scan {
			for _, val := range row {
				if val > 0 && i > lowest {
					lowest = i
				}
			}
		}
		lowest += 2
		sc.drawLine(fmt.Sprintf("0,%d", lowest), fmt.Sprintf("%d,%d", maxSize-1, lowest))
	}

	return sc
}

func (sc *Sandcave) Print() {
	for _, line := range sc.scan {
		fmt.Println(line)
	}
}

func (sc *Sandcave) drawLine(start, stop string) {
	startSplitted := strings.Split(start, positionSeparator)
	stopSplitted := strings.Split(stop, positionSeparator)

	startPos := prepareRockPos(startSplitted)
	stopPos := prepareRockPos(stopSplitted)

	isHorizontal := false
	if startPos.X == stopPos.X {
		isHorizontal = true
	}

	if !isHorizontal {
		if startPos.X <= stopPos.X {
			for col := startPos.X; col <= stopPos.X; col++ {
				sc.scan[startPos.Y][col] = rockVal
			}
		} else {
			for col := startPos.X; col >= stopPos.X; col-- {
				sc.scan[startPos.Y][col] = rockVal
			}
		}
	} else {
		if startPos.Y <= stopPos.Y {
			for row := startPos.Y; row <= stopPos.Y; row++ {
				sc.scan[row][startPos.X] = rockVal
			}
		} else {
			for row := startPos.Y; row >= stopPos.Y; row-- {
				sc.scan[row][startPos.X] = rockVal
			}
		}
	}
}

func prepareRockPos(data []string) *utils.Position {
	pos := &utils.Position{}
	var err error
	pos.X, err = strconv.Atoi(data[0])
	if err != nil {
		log.Fatal(err)
	}

	pos.Y, err = strconv.Atoi(data[1])
	if err != nil {
		log.Fatal(err)
	}

	return pos
}

func newSand() *sand {
	s := &sand{}
	s.X = startX
	s.Y = startY
	s.state = stateInMove
	return s
}

type sand struct {
	utils.Position
	state int
}

func (s *sand) reset() {
	s.X = startX
	s.Y = startY
	s.state = stateInMove
}

func (s *sand) move(sc *Sandcave) {
	if s.Y+1 >= len(sc.scan) || s.X-1 < 0 || s.X+1 >= len(sc.scan[0]) {
		s.state = stateOut
	} else if sc.scan[s.Y+1][s.X] == 0 {
		s.Y += 1
	} else if sc.scan[s.Y+1][s.X-1] == 0 {
		s.Y += 1
		s.X -= 1
	} else if sc.scan[s.Y+1][s.X+1] == 0 {
		s.Y += 1
		s.X += 1
	} else {
		s.state = stateStalled
	}
}

func (sc *Sandcave) CountSand() int {
	counter := 0
	for _, row := range sc.scan {
		for _, val := range row {
			if val == sandVal {
				counter++
			}
		}
	}
	return counter
}
