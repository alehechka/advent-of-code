package twentythree

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func Day16Problem1(inputs []string) string {
	grid := Day16Grid(inputs)

	return strconv.Itoa(grid.CountEnergized())
}

type Day16Grid []string

type Day16Energized map[string]map[rune]int

func (g Day16Grid) CountEnergized() (total int) {
	energized := make(Day16Energized)
	g.fillEnergized(0, 0, Right, energized)
	return len(energized)
}

const (
	Up    rune = '^'
	Down  rune = 'v'
	Left  rune = '<'
	Right rune = '>'
)

func Day16CoordKey(x, y int, direction rune) string {
	return fmt.Sprintf("%d-%d-%s", x, y, string(direction))
}

func (g Day16Grid) fillEnergized(x, y int, direction rune, energized Day16Energized) {
	key := utils.CoordKey(x, y)

	if coord, ok := energized[key]; !ok || coord == nil {
		energized[key] = make(map[rune]int)
	}

	if _, ok := energized[key][direction]; ok {
		return
	}

	energized[key][direction]++

	switch g[y][x] {
	case '.':
		g.attemptToContinue(x, y, direction, energized)
	case '|':
		switch direction {
		case Up, Down:
			g.attemptToContinue(x, y, direction, energized)
		case Right, Left:
			g.attemptToContinue(x, y, Up, energized)
			g.attemptToContinue(x, y, Down, energized)
		}
	case '-':
		switch direction {
		case Up, Down:
			g.attemptToContinue(x, y, Left, energized)
			g.attemptToContinue(x, y, Right, energized)
		case Right, Left:
			g.attemptToContinue(x, y, direction, energized)
		}
	case '/':
		switch direction {
		case Up:
			g.attemptToContinue(x, y, Right, energized)
		case Down:
			g.attemptToContinue(x, y, Left, energized)
		case Left:
			g.attemptToContinue(x, y, Down, energized)
		case Right:
			g.attemptToContinue(x, y, Up, energized)
		}
	case '\\':
		switch direction {
		case Up:
			g.attemptToContinue(x, y, Left, energized)
		case Down:
			g.attemptToContinue(x, y, Right, energized)
		case Left:
			g.attemptToContinue(x, y, Up, energized)
		case Right:
			g.attemptToContinue(x, y, Down, energized)
		}
	}
}

func (g Day16Grid) attemptToContinue(x, y int, direction rune, energized Day16Energized) {
	nextX, nextY, isValid := g.NextCoords(x, y, direction)
	// fmt.Println(x, y, string(g[y][x]), string(direction), nextX, nextX, isValid, energized)
	if isValid {
		g.fillEnergized(nextX, nextY, direction, energized)
	}
}

func (g Day16Grid) NextCoords(x, y int, direction rune) (newX, newY int, isValid bool) {
	switch direction {
	case Up:
		y--
		if y >= 0 {
			isValid = true
		}
	case Down:
		y++
		if y < len(g) {
			isValid = true
		}
	case Left:
		x--
		if x >= 0 {
			isValid = true
		}
	case Right:
		x++
		if x < len(g[y]) {
			isValid = true
		}
	}

	return x, y, isValid
}
