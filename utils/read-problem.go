package utils

import (
	"fmt"
	"os"
	"strings"
)

// ReadProblemInput reads a given day's input by line
func ReadProblemInput(year int, day int) ([]string, error) {
	dat, err := os.ReadFile(fmt.Sprintf("inputs/%d/day-%d.txt", year, day))
	if err != nil {
		return nil, err
	}

	return strings.Split(string(dat), "\n"), nil
}
