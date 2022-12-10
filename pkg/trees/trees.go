package trees

import (
	"fmt"
	"strconv"
)

type Trees struct {
	field          [][]int
	visibility     [][]int
	topScenicScore int
}

func NewTrees(input []string) (*Trees, error) {
	field, err := prepareField(input)
	if err != nil {
		return nil, err
	}
	return &Trees{
		field:          field,
		visibility:     prepareVisibility(input),
		topScenicScore: 0,
	}, nil
}

func (t *Trees) PrintField() {
	printSlice(t.field)
}

func (t *Trees) PrintVisiblility() {
	printSlice(t.visibility)
}

func (t *Trees) CountVisible() int {
	sum := 0
	for _, row := range t.visibility {
		for _, value := range row {
			sum += value
		}
	}
	return sum
}

func (t *Trees) CheckVisibility() {
	for row := 1; row < len(t.field)-1; row++ {
		for col := 1; col < len(t.field[0])-1; col++ {
			topScore, topVisible := t.isTopVisible(row, col)
			bottomScore, bottomVisible := t.isBottomVisible(row, col)
			leftScore, leftVisible := t.isLeftVisible(row, col)
			rightScore, rightVisible := t.isRightVisible(row, col)
			if topVisible || bottomVisible || leftVisible || rightVisible {
				t.visibility[row][col] = 1
			}
			scenicScore := topScore * bottomScore * leftScore * rightScore
			if scenicScore > t.topScenicScore {
				t.topScenicScore = scenicScore
			}
		}
	}
}

func (t *Trees) GetTopScenicScore() int {
	return t.topScenicScore
}

func (t *Trees) isVisible(s []int, row, col int) (int, bool) {
	visible := true
	score := 0
	for _, value := range s {
		score++
		if value >= t.field[row][col] {
			visible = false
			break
		}
	}
	return score, visible
}

func (t *Trees) isTopVisible(row, col int) (int, bool) {
	var tmp []int
	for i := row - 1; i >= 0; i-- {
		tmp = append(tmp, t.field[i][col])
	}
	return t.isVisible(tmp, row, col)
}

func (t *Trees) isBottomVisible(row, col int) (int, bool) {
	var tmp []int
	for i := row + 1; i < len(t.field); i++ {
		tmp = append(tmp, t.field[i][col])
	}
	return t.isVisible(tmp, row, col)

}

func (t *Trees) isLeftVisible(row, col int) (int, bool) {
	var tmp []int
	for i := col - 1; i >= 0; i-- {
		tmp = append(tmp, t.field[row][i])
	}
	return t.isVisible(tmp, row, col)
}

func (t *Trees) isRightVisible(row, col int) (int, bool) {
	var tmp []int
	for i := col + 1; i < len(t.field[0]); i++ {
		tmp = append(tmp, t.field[row][i])
	}
	return t.isVisible(tmp, row, col)
}

func prepareField(input []string) ([][]int, error) {
	field := [][]int{}
	for _, line := range input {
		values := []int{}
		for _, value := range line {
			v, err := strconv.Atoi(string(value))
			if err != nil {
				return nil, err
			}
			values = append(values, v)
		}
		field = append(field, values)
	}
	return field, nil
}

func printSlice(slice [][]int) {
	for _, row := range slice {
		for _, value := range row {
			fmt.Print(value)
		}
		fmt.Println()
	}
}

func prepareVisibility(input []string) [][]int {
	visiblity := [][]int{}
	for row, r := range input {
		var tmp []int
		for col := range r {
			if row == 0 || col == 0 || row == len(r)-1 || col == len(input)-1 {
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, 0)
			}
		}
		visiblity = append(visiblity, tmp)
	}
	return visiblity
}
