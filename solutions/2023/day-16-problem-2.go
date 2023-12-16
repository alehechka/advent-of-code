package twentythree

import (
	"strconv"
	"sync"
)

func Day16Problem2(inputs []string) string {
	grid := Day16Grid(inputs)

	var wg sync.WaitGroup

	var largest int
	for y := range grid {
		wg.Add(1)
		go func(start int) {
			if count := grid.CountEnergizedFromStart(0, start, Right); count > largest {
				largest = count
			}
			if count := grid.CountEnergizedFromStart(len(grid[start])-1, start, Left); count > largest {
				largest = count
			}
			wg.Done()
		}(y)
	}

	for x := range grid[0] {
		wg.Add(1)
		go func(start int) {
			if count := grid.CountEnergizedFromStart(start, 0, Down); count > largest {
				largest = count
			}
			if count := grid.CountEnergizedFromStart(start, len(grid)-1, Up); count > largest {
				largest = count
			}
			wg.Done()
		}(x)
	}

	wg.Wait()

	return strconv.Itoa(largest)
}

func (g Day16Grid) CountEnergizedFromStart(x, y int, direction rune) (total int) {
	energized := make(Day16Energized)
	g.fillEnergized(x, y, direction, energized)
	return len(energized)
}
