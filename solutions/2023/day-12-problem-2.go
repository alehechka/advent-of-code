package twentythree

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func Day12Problem2(inputs []string) string {
	rows := Day12ParseInputUnfolded(inputs)

	var sum int
	// fmt.Println(rows[len(rows)-1].MatchesTotal())
	for _, row := range rows {
		sum += row.MatchesTotal()
	}

	return strconv.Itoa(sum)
}

func (row Day12Row) MatchesTotal() int {
	states := "."
	for _, num := range row.StatusGroups {
		states += strings.Repeat("#", num)
		states += "."
	}

	states_dict := map[int]int{0: 1}
	new_dict := make(map[int]int)

	for _, ch := range row.Springs {
		for state := range states_dict {
			switch ch {
			case '?':
				if state+1 < len(states) {
					new_dict[state+1] = new_dict[state+1] + states_dict[state]
				}
				if states[state] == '.' {
					new_dict[state] = new_dict[state] + states_dict[state]
				}
			case '.':
				if state+1 < len(states) && states[state+1] == '.' {
					new_dict[state+1] = new_dict[state+1] + states_dict[state]
				}
				if states[state] == '.' {
					new_dict[state] = new_dict[state] + states_dict[state]
				}
			case '#':
				if state+1 < len(states) && states[state+1] == '#' {
					new_dict[state+1] = new_dict[state+1] + states_dict[state]
				}
			}
		}

		states_dict = new_dict
		new_dict = make(map[int]int)
	}

	sum := states_dict[len(states)-1] + states_dict[len(states)-2]
	return sum
}

func Day12Problem2_BruteForce(inputs []string) string {
	rows := Day12ParseInputUnfolded(inputs)

	var sum int
	var wg sync.WaitGroup
	for _, row := range rows {
		wg.Add(1)
		go func(r Day12Row) {
			count := r.FindRecursivePossibilities()
			fmt.Println(r, count)
			sum += count
			wg.Done()
		}(row)
	}

	wg.Wait()

	return strconv.Itoa(sum)
}

func (r Day12Row) FindRecursivePossibilities() int {
	return r.findRecursivePossibilities(r.Springs, '.', 0, 0, "")
}

func (r Day12Row) findRecursivePossibilities(str string, prevChar byte, groupIndex int, groupCount int, builder string) (count int) {
	// fmt.Println(str, string(prevChar), groupIndex, groupCount, builder)

	if groupIndex >= len(r.StatusGroups) {
		if len(str) == 0 {
			// fmt.Printf("%s => %s %v\n", r.Springs, builder, r.StatusGroups)
			return 1
		}
		if len(str) > 0 && (str[0] == '.' || str[0] == '?') {
			return r.findRecursivePossibilities(str[1:], '.', groupIndex, groupCount, builder+string(str[0]))
		}
		return 0
	}

	if len(str) == 0 {
		return 0
	}

	expectedGroupCount := r.StatusGroups[groupIndex]
	if groupCount > expectedGroupCount {
		return 0
	}

	ch := str[0]
	hasMoreChars := len(str) > 1
	isLastGroup := groupIndex == len(r.StatusGroups)-1

	if ch == '.' || ch == '#' {
		builder += string(ch)
	}

	if ch == '.' {
		if hasMoreChars {
			if prevChar == '#' {
				if groupCount != expectedGroupCount {
					return 0
				}
				groupIndex++
			}

			return r.findRecursivePossibilities(str[1:], ch, groupIndex, 0, builder)
		}

		if isLastGroup && groupCount == expectedGroupCount {
			// fmt.Printf("%s => %s %v\n", r.Springs, builder, r.StatusGroups)
			return 1
		}

		return 0
	}

	if ch == '#' {
		groupCount++
		if hasMoreChars {
			return r.findRecursivePossibilities(str[1:], ch, groupIndex, groupCount, builder)
		}

		if isLastGroup && groupCount == expectedGroupCount {
			// fmt.Printf("%s => %s %v\n", r.Springs, builder, r.StatusGroups)
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
