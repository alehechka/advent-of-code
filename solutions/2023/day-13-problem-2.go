package twentythree

import (
	"strconv"
)

func Day13Problem2(inputs []string) string {
	patterns := Day13ExtractPatterns(inputs)

	var sum int
	for _, pattern := range patterns {
		sum += pattern.FindSmudgedPatternValue()
	}

	return strconv.Itoa(sum)
}

func (p Day13Pattern) FindSmudgedPatternValue() (value int) {
	ogValue := p.FindPatternValue()
	skipX, skipY := -1, -1
	if ogValue >= 100 {
		skipX = ogValue / 100
	} else {
		skipY = ogValue
	}

	for y, row := range p {
		for x, ch := range row {
			copy := p.DeepCopy()

			newChar := '#'
			if ch == '#' {
				newChar = '.'
			}

			newRow := ""
			if x > 0 {
				newRow += row[0:x]
			}
			newRow += string(newChar)
			if x+1 < len(row) {
				newRow += row[x+1:]
			}
			copy[y] = newRow

			value := copy.FindPatternValueWithSkip(skipX, skipY)

			if value > 0 && value != ogValue {
				return value
			}
		}
	}
	return
}

func (p Day13Pattern) FindPatternValueWithSkip(skipX int, skipY int) int {
	if value := p.FindHorizontalMirrorIndexWithSkip(skipX); value > 0 {
		return value * 100
	}
	return p.FindVerticalMirrorIndexWithSkip(skipY)
}

func (p Day13Pattern) FindVerticalMirrorIndexWithSkip(skip int) (mirror int) {

	for index := range p[0] {
		if index != skip && p.IsVerticalValid(index) {
			return index
		}
	}

	return
}

func (p Day13Pattern) FindHorizontalMirrorIndexWithSkip(skip int) (mirror int) {

	for index := range p {
		if index != skip && p.IsHorizontalValid(index) {
			return index
		}
	}

	return
}

func (p Day13Pattern) DeepCopy() (output Day13Pattern) {
	for index := range p {
		output = append(output, p[index])
	}
	return
}
