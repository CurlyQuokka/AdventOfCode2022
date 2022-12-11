package video

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	opSeparator  = " "
	darkPixel    = "."
	litPixel     = "#"
	screenWidth  = 40
	screenHeight = 6
)

type Video struct {
	register   int
	clock      int
	signal     int
	finished   chan (bool)
	limit      int
	checkpoint int
	increase   int
	screen     [screenWidth * screenHeight]string
}

func NewVideo() *Video {
	return &Video{
		register:   1,
		clock:      0,
		signal:     0,
		limit:      220,
		checkpoint: 20,
		increase:   screenWidth,
		screen:     [screenWidth * screenHeight]string{},
	}
}

func (v *Video) incremetnClock() {
	if v.clock >= v.register-1+(v.clock-v.clock%screenWidth) && v.clock <= v.register+1+(v.clock-v.clock%screenWidth) {
		v.screen[v.clock] = litPixel
	}
	v.clock++
	if v.clock == v.checkpoint {
		v.signal += v.register * v.clock
		if v.checkpoint < v.limit {
			v.checkpoint += v.increase
		}
	}
}

func (v *Video) addOp(value int) {
	for i := 0; i < 2; i++ {
		v.incremetnClock()
	}
	v.register += value
}

func (v *Video) noOp() {
	v.incremetnClock()
}

func (v *Video) ProcessInput(input []string) error {
	for i := 0; i < len(v.screen); i++ {
		v.screen[i] = darkPixel
	}
	for _, op := range input {
		opSplitted := strings.Split(op, opSeparator)
		if len(opSplitted) < 2 {
			v.noOp()
		} else {
			value, err := strconv.Atoi(opSplitted[1])
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			v.addOp(value)
		}
	}
	return nil
}

func (v *Video) GetSignalStrength() int {
	return v.signal
}

func (v *Video) ShowScreen() {
	for i := 0; i < len(v.screen); i++ {
		if i%screenWidth == 0 {
			fmt.Println()
		}
		fmt.Print(v.screen[i])
	}
	fmt.Println()
}
