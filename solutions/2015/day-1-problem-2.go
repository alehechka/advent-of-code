package fifteen

import "strconv"

func Day1Problem2(inputs []string) string {
	var floor int

	for index, ch := range inputs[0] {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}
		if floor < 0 {
			return strconv.Itoa(index + 1)
		}
	}

	return ""
}
