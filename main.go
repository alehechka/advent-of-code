package main

import (
	"advent/solutions"
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments supplied, please provide year, day, and problem numbers")
		return
	}

	solution, year, day, _, err := FindSolution()
	if err != nil {
		fmt.Println(err)
		return
	}

	problemInput, err := utils.ReadProblemInput(year, day)
	if err != nil {
		fmt.Println(err)
		return
	}

	start := time.Now()
	result := solution(problemInput)
	fmt.Printf("Result: %s\n", result)
	fmt.Printf("Duration: %v\n", time.Since(start))
}

func FindSolution() (solution solutions.Solution, year, day, problem int, err error) {
	year, err = strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	solutionFuncs, ok := solutions.Solutions[year]
	if !ok {
		keys := make([]int, 0)
		for k := range solutions.Solutions {
			keys = append(keys, k)
		}
		err = fmt.Errorf("year %d is not valid, options are: %v", year, keys)
		return
	}

	day, err = strconv.Atoi(os.Args[2])
	if err != nil {
		return
	}
	if day < 1 || day > 25 {
		err = fmt.Errorf("day %d is invalid, please provide between 1 and 25", day)
		return
	}
	if day > len(solutionFuncs) {
		err = fmt.Errorf("day %d is not ready yet, please choose another", day)
		return
	}

	problem, err = strconv.Atoi(os.Args[3])
	if err != nil {
		return
	}

	if problem < 1 || problem > 2 {
		err = fmt.Errorf("problem %d is invalid, please provide 1 or 2", problem)
		return
	}

	if problem > len(solutionFuncs[day-1]) {
		err = fmt.Errorf("problem %d is not ready yet, please choose another", problem)
		return
	}

	solution = solutionFuncs[day-1][problem-1]
	return
}
