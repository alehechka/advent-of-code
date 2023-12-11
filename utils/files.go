package utils

import "os"

func WriteMatrixToFile(pipes [][]rune, fileName string) {
	var data []byte
	for index, line := range pipes {
		for _, ch := range line {
			data = append(data, byte(ch))
		}
		if index < len(pipes)-1 {
			data = append(data, '\n')
		}
	}
	if err := os.WriteFile(fileName, data, 0644); err != nil {
		panic(err)
	}
}
