package utils

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// ReadProblemInput reads a given day's input by line
func ReadProblemInput(day int, problem int) ([]string, error) {
	dat, err := os.ReadFile(path.Join("inputs", fmt.Sprintf("day-%d", day), fmt.Sprintf("problem-%d.txt", problem)))
	if err != nil {
		return nil, err
	}

	return strings.Split(string(dat), "\n"), nil
}
