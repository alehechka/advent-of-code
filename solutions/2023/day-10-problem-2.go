package twentythree

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func Day10Problem2(inputs []string) string {
	pipeCoords, _ := Day10TraverseLoop(inputs)

	return strconv.Itoa(Day10CountInclosed(Day10LoopOnly(inputs, pipeCoords), pipeCoords))
}

func Day10CountInclosed(rows [][]rune, pipeCoords map[string]int) (count int) {
	loop := make([][]rune, len(rows))

	for y, row := range rows {
		if loop[y] == nil {
			loop[y] = make([]rune, len(rows[y]))
		}
		for x, ch := range row {
			if _, ok := pipeCoords[Day10Key(x, y)]; ok {
				loop[y][x] = ch
			} else if Day10IsInner(rows, x, y) {
				count++
				loop[y][x] = '.'
			} else {
				loop[y][x] = ' '
			}
		}
	}
	utils.WriteMatrixToFile(loop, "day-10-next.out")
	return
}

func Day10IsInner(rows [][]rune, x, y int) bool {
	row := rows[y]
	fmt.Printf("Checking: %s, at %dx%d\n", string(row[x]), x, y)

	var left, right, above, below bool

	for i := 1; i < len(rows); i++ {
		if x-i >= 0 && row[x-i] != '.' && !left {
			left = true
			fmt.Printf("Crossed Left: %s, at %dx%d\n", string(row[x-1]), x-i, y)
		}
		if x+i+1 < len(row) && row[x+i] != '.' && !right {
			right = true
			fmt.Printf("Crossed Right: %s, at %dx%d\n", string(row[x+i]), x+i, y)
		}
		if y-i >= 0 && rows[y-i][x] != '.' && !above {
			above = true
			fmt.Printf("Crossed Above: %s, at %dx%d\n", string(rows[y-i][x]), x, y-i)
		}
		if y+i+1 < len(rows) && rows[y+i][x] != '.' && !below {
			below = true
			fmt.Printf("Crossed Below: %s, at %dx%d\n", string(rows[y+i][x]), x, y+i)
		}
		if left && right && above && below {
			return true
		}
	}

	return left && right && above && below
}

func Day10LoopOnly(pipes []string, pipeCoords map[string]int) (loop [][]rune) {
	loop = make([][]rune, len(pipes))

	for y, pipe := range pipes {
		if loop[y] == nil {
			loop[y] = make([]rune, len(pipes[y]))
		}
		for x, ch := range pipe {
			if count, ok := pipeCoords[Day10Key(x, y)]; ok && count > 0 {
				loop[y][x] = ch
			} else {
				loop[y][x] = '.'
			}
		}
	}

	return
}
