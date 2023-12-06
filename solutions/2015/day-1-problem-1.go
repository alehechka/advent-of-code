package fifteen

import "strconv"

func Day1Problem1(inputs []string) string {
	var floor int

	for _, ch := range inputs[0] {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}
	}

	return strconv.Itoa(floor)
}
