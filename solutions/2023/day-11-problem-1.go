package twentythree

import "strconv"

func Day11Problem1(inputs []string) string {
	universe := Day11ExpandUniverse(inputs)

	galaxies := Day11FindGalaxies(universe)

	return strconv.Itoa(galaxies.ShortestDistanceSum())
}

type Day11Coord struct {
	X int
	Y int
}

func (c Day11Coord) DistanceTo(other Day11Coord) (distance int) {
	return Abs(c.X-other.X) + Abs(c.Y-other.Y)
}

func Abs(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}

type Day11Coords []Day11Coord

func (coords Day11Coords) ShortestDistanceSum() (sum int) {
	for i, coord := range coords {
		for i2 := i + 1; i2 < len(coords); i2++ {
			sum += coord.DistanceTo(coords[i2])
		}
	}

	return
}

func Day11FindGalaxies(universe [][]rune) Day11Coords {
	galaxies := make([]Day11Coord, 0)

	for y, row := range universe {
		for x, ch := range row {
			if ch == '#' {
				galaxies = append(galaxies, Day11Coord{X: x, Y: y})
			}
		}
	}

	return galaxies
}

func Day11ExpandUniverse(inputs []string) (universe [][]rune) {
	verticalGalaxies := make(map[int]bool)
	horizontalGalaxies := make(map[int]bool)

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

	tempUniverse := make([][]rune, 0)
	for x := range inputs[0] {
		hasGalaxy, ok := verticalGalaxies[x]
		for y := range inputs {
			if len(tempUniverse) < y+1 {
				tempUniverse = append(tempUniverse, make([]rune, 0))
			}

			tempUniverse[y] = append(tempUniverse[y], rune(inputs[y][x]))
			if !ok || !hasGalaxy {
				tempUniverse[y] = append(tempUniverse[y], '.')
			}
		}
	}

	for y := range inputs {
		universe = append(universe, tempUniverse[y])
		if hasGalaxy, ok := horizontalGalaxies[y]; !ok || !hasGalaxy {
			universe = append(universe, day11CreateEmptyUniverseRow(len(tempUniverse[y])))
		}
	}

	return
}

func day11CreateEmptyUniverseRow(length int) []rune {
	row := make([]rune, length)

	for index := range row {
		row[index] = '.'
	}

	return row
}
