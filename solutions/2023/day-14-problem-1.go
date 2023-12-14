package twentythree

import (
	"strconv"
	"strings"
)

func Day14Problem1(inputs []string) string {
	platform := Day14Platform(inputs)

	loaded := platform.LoadNorth()

	return strconv.Itoa(loaded.CalculateNorthLoad())
}

type Day14Platform []string

func (p Day14Platform) LoadNorth() Day14Platform {
	return p.loadVertical(Day14ShiftLeft)
}

func (p Day14Platform) loadVertical(loader func(string) string) (newPlatform Day14Platform) {
	columns := make([]string, 0)

	for x := range p[0] {
		var column string
		for y := range p {
			column += string(p[y][x])
		}
		columns = append(columns, column)
	}

	newPlatform = make(Day14Platform, len(p))
	for _, column := range columns {
		leftShifted := loader(column)
		for index, ch := range leftShifted {
			newPlatform[index] += string(ch)
		}
	}
	return
}

func Day14ShiftLeft(str string) (result string) {
	var spaceCount int
	for _, ch := range str {
		if ch == '.' {
			spaceCount++
		}
		if ch == '#' {
			result += strings.Repeat(".", spaceCount)
			spaceCount = 0
		}
		if ch == '#' || ch == 'O' {
			result += string(ch)
		}

	}

	result += strings.Repeat(".", spaceCount)

	return result
}

func (p Day14Platform) CalculateNorthLoad() (load int) {
	for rowIndex, row := range p {
		for _, ch := range row {
			if ch == 'O' {
				load += len(p) - rowIndex
			}
		}
	}

	return
}
