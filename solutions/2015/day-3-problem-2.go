package fifteen

import (
	"strconv"
)

func Day3Problem2(inputs []string) string {
	var x1, y1, x2, y2 int
	houses := map[string]int{Day3Key(x1, y1): 2}

	for index, direction := range inputs[0] {
		if index%2 == 0 {
			x1, y1 = Day3Increment(x1, y1, direction)
			houses[Day3Key(x1, y1)]++
		} else {
			x2, y2 = Day3Increment(x2, y2, direction)
			houses[Day3Key(x2, y2)]++
		}
	}

	return strconv.Itoa(len(houses))
}

func Day3Increment(x, y int, d rune) (int, int) {
	switch d {
	case '^':
		return x, y + 1
	case 'v':
		return x, y - 1
	case '>':
		return x + 1, y
	case '<':
		return x - 1, y
	default:
		return x, y
	}
}
