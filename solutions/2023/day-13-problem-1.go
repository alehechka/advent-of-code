package twentythree

import (
	"strconv"
)

func Day13Problem1(inputs []string) string {
	patterns := Day13ExtractPatterns(inputs)

	var sum int
	for _, pattern := range patterns {
		sum += pattern.FindPatternValue()
	}

	return strconv.Itoa(sum)
}

type Day13Pattern []string

func (p Day13Pattern) FindPatternValue() int {
	if value := p.FindHorizontalMirrorIndex(); value > 0 {
		return value * 100
	}
	return p.FindVerticalMirrorIndex()
}

func (p Day13Pattern) FindVerticalMirrorIndex() (mirror int) {

	for index := range p[0] {
		if p.IsVerticalValid(index) {
			return index
		}
	}

	return
}

func (p Day13Pattern) IsVerticalValid(index int) bool {
	if index == 0 {
		return false
	}

	for _, row := range p {
		for i := 0; i+index < len(row) && index-i-1 >= 0; i++ {
			if row[i+index] != row[index-i-1] {
				return false
			}
		}
	}

	return true
}

func (p Day13Pattern) FindHorizontalMirrorIndex() (mirror int) {

	for index := range p {
		if p.IsHorizontalValid(index) {
			return index
		}
	}

	return
}

func (p Day13Pattern) IsHorizontalValid(index int) bool {
	if index == 0 {
		return false
	}

	for i := 0; i+index < len(p) && index-i-1 >= 0; i++ {
		if p[i+index] != p[index-i-1] {
			return false
		}
	}
	return true
}

func Day13ExtractPatterns(inputs []string) (patterns []Day13Pattern) {
	var pattern Day13Pattern
	for _, input := range inputs {
		if input != "" {
			pattern = append(pattern, input)
		} else {
			patterns = append(patterns, pattern)
			pattern = make(Day13Pattern, 0)
		}
	}

	if len(pattern) > 0 {
		patterns = append(patterns, pattern)
	}

	return
}
