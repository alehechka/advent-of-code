package twentythree

import (
	"strconv"
	"strings"
)

func Day12Problem2(inputs []string) string {
	rows := Day12ParseInput(inputs)

	var sum int
	for _, row := range rows {
		count := row.FindRecursivePossibilities()
		sum += count
	}

	return strconv.Itoa(sum)
}

func (r Day12Row) FindRecursivePossibilities() int {
	return r.findRecursivePossibilities(r.Springs, '.', 0, 0, "")
}

func (r Day12Row) findRecursivePossibilities(str string, prevChar byte, groupIndex int, groupCount int, builder string) (count int) {
	if len(str) == 0 {
		return 0
	}

	if groupIndex >= len(r.StatusGroups) {
		return 0
	}

	expectedGroupCount := r.StatusGroups[groupIndex]
	if groupCount > expectedGroupCount {
		return 0
	}

	ch := str[0]
	hasMoreChars := len(str) > 1
	isLastGroup := groupIndex == len(r.StatusGroups)-1

	if ch == '.' {
		if hasMoreChars {
			if prevChar == '#' {
				groupIndex++
			}

			return r.findRecursivePossibilities(str[1:], ch, groupIndex, 0, builder+string(ch))
		}

		if isLastGroup && groupCount == expectedGroupCount {
			// fmt.Printf("%s => %s%s %v\n", r.Springs, builder, string(ch), r.StatusGroups)
			return 1
		}

		return 0
	}

	if ch == '#' {
		groupCount++
		if hasMoreChars {
			return r.findRecursivePossibilities(str[1:], ch, groupIndex, groupCount, builder+string(ch))
		}

		if isLastGroup && groupCount == expectedGroupCount {
			// fmt.Printf("%s => %s%s %v\n", r.Springs, builder, string(ch), r.StatusGroups)
			return 1
		}

		return 0
	}

	if ch == '?' {
		return r.findRecursivePossibilities("."+str[1:], prevChar, groupIndex, groupCount, builder) +
			r.findRecursivePossibilities("#"+str[1:], prevChar, groupIndex, groupCount, builder)
	}

	return
}

func Day12ParseInputUnfolded(inputs []string) (rows []Day12Row) {
	for _, input := range inputs {
		rows = append(rows, Day12ParseRowUnfolded(input))
	}
	return
}

func Day12ParseRowUnfolded(input string) (row Day12Row) {
	parts := strings.Split(input, " ")

	row.Springs = strings.Join([]string{parts[0], parts[0], parts[0], parts[0], parts[0]}, "?")

	nums := strings.Split(parts[1], ",")
	for i := 0; i < 5; i++ {
		for _, num := range nums {
			if x, err := strconv.Atoi(num); err == nil {
				row.StatusGroups = append(row.StatusGroups, x)
			}
		}
	}

	return
}
