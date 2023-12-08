package fifteen

import (
	"fmt"
	"strconv"
)

func Day3Problem1(inputs []string) string {
	var x, y int
	houses := map[string]int{Day3Key(x, y): 1}

	for _, direction := range inputs[0] {
		x, y = Day3Increment(x, y, direction)
		houses[Day3Key(x, y)]++
	}

	return strconv.Itoa(len(houses))
}

func Day3Key(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}
