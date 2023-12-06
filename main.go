package main

import (
	"advent/solutions"
	"advent/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments supplied, please provide year, day, and problem numbers")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	solutionFuncs, ok := solutions.Solutions[year]
	if !ok {
		keys := make([]int, 0)
		for k := range solutions.Solutions {
			keys = append(keys, k)
		}
		fmt.Printf("Year %d is not valid, options are: %v\n", year, keys)
		return
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	if day < 1 || day > 25 {
		fmt.Printf("Day %d is invalid, please provide between 1 and 25\n", day)
		return
	}
	if day > len(solutionFuncs) {
		fmt.Printf("Day %d is not ready yet, please choose another\n", day)
		return
	}

	problem, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}

	if problem < 1 || problem > 2 {
		fmt.Printf("Problem %d is invalid, please provide 1 or 2\n", problem)
		return
	}

	if problem > len(solutionFuncs[day-1]) {
		fmt.Printf("Problem %d is not ready yet, please choose another\n", problem)
		return
	}

	problemInput, err := utils.ReadProblemInput(year, day)
	if err != nil {
		panic(err)
	}

	result := solutionFuncs[day-1][problem-1](problemInput)
	fmt.Printf("Result: %s\n", result)
}
