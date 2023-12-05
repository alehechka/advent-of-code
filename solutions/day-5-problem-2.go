package solutions

import (
	"fmt"
	"strconv"
)

func Day5Problem2(inputs []string) string {
	seeds := Day5ParseSeeds(inputs[0])

	categoryMap := Day5ParseCategoryMap(inputs[1:])

	var lowest int
	for index, seed := range seeds {
		if index%2 != 0 {
			continue
		}
		fmt.Printf("Traversing from %d to %d, %d items\n", seed, seed+seeds[index+1], seeds[index+1])
		for i := 0; i <= seeds[index+1]; i++ {
			dest := categoryMap.Traverse("seed", "location", seed+i)
			if lowest == 0 || dest < lowest {
				lowest = dest
			}
		}
	}

	return strconv.Itoa(lowest)
}
