package twentythree

import (
	"strconv"
)

func Day11Problem2(inputs []string) string {
	verticalGalaxies, horizontalGalaxies := Day11FindExpansionPoints(inputs)

	galaxies := Day11FindGalaxiesWithExpansion(inputs, verticalGalaxies, horizontalGalaxies, 1000000)

	return strconv.Itoa(galaxies.ShortestDistanceSum())
}

type Day11GalaxyCoords map[int]bool

func (c Day11GalaxyCoords) IsGalaxy(coord int) bool {
	hasGalaxy, ok := c[coord]
	return ok && hasGalaxy
}

func Day11FindGalaxiesWithExpansion(universe []string, vertical, horizontal Day11GalaxyCoords, expansion int) Day11Coords {
	galaxies := make([]Day11Coord, 0)

	var verticalExpansion int
	for y, row := range universe {
		if !horizontal.IsGalaxy(y) {
			verticalExpansion += expansion - 1
		}
		var horizontalExpansion int
		for x, ch := range row {
			if !vertical.IsGalaxy(x) {
				horizontalExpansion += expansion - 1
			}
			if ch == '#' {
				galaxies = append(galaxies, Day11Coord{X: x + horizontalExpansion, Y: y + verticalExpansion})
			}
		}
	}

	return galaxies
}

func Day11FindExpansionPoints(inputs []string) (verticalGalaxies, horizontalGalaxies map[int]bool) {
	verticalGalaxies = make(map[int]bool)
	horizontalGalaxies = make(map[int]bool)

	for y, input := range inputs {
		var hasHorizontalGalaxy bool
		for x, ch := range input {
			if ch == '#' {
				hasHorizontalGalaxy = true
				verticalGalaxies[x] = true
			}
		}

		horizontalGalaxies[y] = hasHorizontalGalaxy
	}

	return
}
