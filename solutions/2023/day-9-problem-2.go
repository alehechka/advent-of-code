package twentythree

import "strconv"

func Day9Problem2(inputs []string) string {
	var total int
	for _, input := range inputs {
		pattern := Day9Pattern(ParseInts(input))
		next := pattern.Previous()
		total += next
	}

	return strconv.Itoa(total)
}

func (p Day9Pattern) Previous() int {
	nextPattern := make(Day9Pattern, 0)
	var numZeroes int

	for index := range p {
		if index+1 < len(p) {
			difference := p[index+1] - p[index]
			nextPattern = append(nextPattern, difference)
			if difference == 0 {
				numZeroes++
			}
		}
	}

	if numZeroes == len(p)-1 {
		return p[0]
	}

	return p[0] - nextPattern.Previous()
}
