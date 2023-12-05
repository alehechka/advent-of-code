package solutions

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Day5Problem2(inputs []string) string {
	seeds := Day5ParseSeeds(inputs[0])

	categoryMap := Day5ParseCategoryMap(inputs[1:])

	var lowest int
	var wg sync.WaitGroup
	for index, seed := range seeds {
		if index%2 != 0 {
			continue
		}
		wg.Add(1)
		go func(index, seed int) {
			dest := categoryMap.TraverseRange("seed", "location", seed, seeds[index+1])
			if lowest == 0 || dest < lowest {
				lowest = dest
			}
			wg.Done()
		}(index, seed)
	}

	wg.Wait()

	return strconv.Itoa(lowest)
}

func (c Day5CategoryMap) TraverseRange(startCategory string, endCategory string, startIndex int, seedRange int) int {
	fmt.Printf("Traversing from %d to %d, %d items\n", startIndex, startIndex+seedRange, seedRange)
	startTime := time.Now()

	var lowest int
	for i := 0; i <= seedRange; i++ {
		dest := c.Traverse("seed", "location", startIndex+i)
		if lowest == 0 || dest < lowest {
			lowest = dest
		}
	}

	fmt.Printf("Traversed from %d to %d, %d items in %v, found lowest: %d\n", startIndex, startIndex+seedRange, seedRange, time.Since(startTime), lowest)
	return lowest
}
