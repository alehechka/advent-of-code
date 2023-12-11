package twentythree

import (
	"os"
	"strconv"
)

func Day10Problem2(inputs []string) string {
	pipeCoords, _ := Day10TraverseLoop(inputs)

	return strconv.Itoa(Day10CountInclosed(inputs, pipeCoords))
}

func Day10CountInclosed(rows []string, pipeCoords map[string]int) (count int) {
	loop := make([][]rune, len(rows))

	for y, row := range rows {
		if loop[y] == nil {
			loop[y] = make([]rune, len(rows[y]))
		}
		for x, ch := range row {
			if _, ok := pipeCoords[Day10Key(x, y)]; !ok {
				if Day10IsInner(rows, x, y, pipeCoords) {
					count++
					loop[y][x] = '.'
				} else {
					loop[y][x] = ' '
				}
			} else {
				loop[y][x] = ch
			}
		}
	}
	Day10WriteLoopOnly(loop, "day-10-next.out")
	return
}

func Day10IsInner(rows []string, x, y int, pipeCoords map[string]int) bool {
	var crossCount int

	for {
		x--
		y--
		if x < 0 || y < 0 {
			break
		}
		if _, ok := pipeCoords[Day10Key(x, y)]; ok {
			crossCount++
		}
	}

	return crossCount > 0 && crossCount%2 != 0
}

func Day10WriteLoopOnly(pipes [][]rune, fileName string) {
	var data []byte
	for _, line := range pipes {
		for _, ch := range line {
			data = append(data, byte(ch))
		}
		data = append(data, '\n')
	}
	if err := os.WriteFile(fileName, data, 0644); err != nil {
		panic(err)
	}
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
