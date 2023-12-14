package twentythree

import (
	"strconv"
	"strings"
)

func Day14Problem2(inputs []string) string {
	keys := make(map[string]int)

	platform := Day14Platform(inputs)

	keys[platform.Key()] = 0

	cycles := 1_000_000_000
	foundLoop := false
	for i := 1; i <= cycles; i++ {
		platform = platform.SpinCycle()

		if !foundLoop {
			key := platform.Key()
			if prev, ok := keys[key]; ok {
				foundLoop = true
				loopLength := i - prev
				remaining := cycles - ((cycles - prev) % loopLength)
				i = remaining
			}
			keys[key] = i
		}
	}

	return strconv.Itoa(platform.CalculateNorthLoad())
}

func (p Day14Platform) Key() string {
	return strings.Join(p, "")
}

func (p Day14Platform) SpinCycle() (newPlatform Day14Platform) {
	north := p.LoadNorth()
	west := north.LoadWest()
	south := west.LoadSouth()
	east := south.LoadEast()

	return east
}

func (p Day14Platform) LoadSouth() Day14Platform {
	return p.loadVertical(Day14ShiftRight)
}

func (p Day14Platform) LoadEast() Day14Platform {
	return p.loadHorizontal(Day14ShiftRight)
}

func (p Day14Platform) LoadWest() Day14Platform {
	return p.loadHorizontal(Day14ShiftLeft)
}

func (p Day14Platform) loadHorizontal(loader func(string) string) (newPlatform Day14Platform) {
	for _, row := range p {
		newPlatform = append(newPlatform, loader(row))
	}
	return
}

func Day14ShiftRight(str string) (result string) {
	var spaceCount int
	for i := len(str) - 1; i >= 0; i-- {
		ch := str[i]
		if ch == '.' {
			spaceCount++
		}
		if ch == '#' {
			result = strings.Repeat(".", spaceCount) + result
			spaceCount = 0
		}
		if ch == '#' || ch == 'O' {
			result = string(ch) + result
		}

	}

	result = strings.Repeat(".", spaceCount) + result

	return result
}
