package twentythree

import (
	"strconv"
	"strings"
)

func Day15Problem1(inputs []string) string {
	codes := strings.Split(inputs[0], ",")

	var sum int
	for _, code := range codes {
		sum += Day15HASH(code)
	}

	return strconv.Itoa(sum)
}

func Day15HASH(str string) (value int) {
	for _, ch := range str {
		value += int(ch)
		value *= 17
		value %= 256
	}
	return
}
