package communication

func ProcessSignal(signal string, windowSize int) int {
	for i := windowSize - 1; i < len(signal); i++ {
		frame := signal[i-(windowSize-1) : i+1]
		unique := getNumOfUnique(frame)
		if unique == windowSize {
			return i + 1
		}
	}
	return 0
}

func getNumOfUnique(frame string) int {
	cMap := make(map[rune]bool)
	for _, r := range frame {
		cMap[r] = true
	}
	return len(cMap)
}
