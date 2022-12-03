package supplies

const (
	lowercase     = 96
	uppercase     = 64
	numberOfChars = 26
	groupSize     = 3
)

type supplies struct {
	rucksacks []*rucksack
	doubled   []rune
}

type rucksack struct {
	conpartments []string
}

func newRucksack(in string) *rucksack {
	r := &rucksack{}
	halfIndex := len(in) / 2
	r.conpartments = append(r.conpartments, in[:halfIndex])
	r.conpartments = append(r.conpartments, in[halfIndex:])
	return r
}

func NewSupplies(input []string) *supplies {
	s := &supplies{
		rucksacks: []*rucksack{},
	}
	for _, line := range input {
		s.rucksacks = append(s.rucksacks, newRucksack(line))
	}
	return s
}

func (r *rucksack) findDoubled() []rune {
	doubled := []rune{}
	itemMap := make(map[rune]int)
	for _, item := range r.conpartments[0] {
		itemMap[item] = 1
	}

	for _, item := range r.conpartments[1] {
		if _, exists := itemMap[item]; exists {
			doubled = append(doubled, item)
			break
		}
	}

	return doubled
}

func (s *supplies) findDoubled() []rune {
	doubled := []rune{}
	for _, r := range s.rucksacks {
		doubled = append(doubled, r.findDoubled()...)
	}
	return doubled
}

func getPriority(r rune) uint8 {
	if r >= lowercase {
		return uint8(r) - lowercase
	}
	return uint8(r) - uppercase + numberOfChars
}

func (s *supplies) GetSumOfPriorites() uint {
	doubled := s.findDoubled()
	var sum uint
	sum = 0
	for _, r := range doubled {
		sum += uint(getPriority(r))
	}
	return sum
}

func (s *supplies) GetSumOfGroupPriorities() int {
	groupItems := []rune{}
	for i := 0; i < len(s.rucksacks); i = i + groupSize {
		groupItems = append(groupItems, s.findGroupItem(i))
	}
	sum := 0
	for _, r := range groupItems {
		sum += int(getPriority(r))
	}
	return sum
}

func (s *supplies) findGroupItem(index int) rune {
	r := []string{}
	for i := 0; i < groupSize; i++ {

		r = append(r, removeMulti(s.rucksacks[index+i].conpartments[0]+s.rucksacks[index+i].conpartments[1]))
	}
	itemMap := make(map[rune]uint)

	for i, rs := range r {
		for _, item := range rs {
			if _, exists := itemMap[item]; exists {
				if itemMap[item] < uint(i+1) {
					itemMap[item]++
				}
			} else {
				itemMap[item] = 1
			}
		}
	}

	for key, value := range itemMap {
		if value == groupSize {
			return key
		}
	}

	return 0
}

func removeMulti(s string) string {
	sMap := make(map[rune]bool)
	for _, r := range s {
		sMap[r] = true
	}
	out := ""
	for key, _ := range sMap {
		out += string(key)
	}
	return out
}
