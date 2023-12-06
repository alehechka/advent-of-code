package main

import (
	"advent/solutions"
	"advent/utils"
	"fmt"
	"os"
	"strconv"
)

var solutionFuncs = [][]func([]string) string{
	{solutions.Day1Problem1, solutions.Day1Problem2},
	{solutions.Day2Problem1, solutions.Day2Problem2},
	{solutions.Day3Problem1, solutions.Day3Problem2},
	{solutions.Day4Problem1, solutions.Day4Problem2},
	{solutions.Day5Problem1, solutions.Day5Problem2},
	{solutions.Day6Problem1},
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments supplied, please provide day and problem numbers")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
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

	problem, err := strconv.Atoi(os.Args[2])
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

	problemInput, err := utils.ReadProblemInput(day)
	if err != nil {
		panic(err)
	}

	result := solutionFuncs[day-1][problem-1](problemInput)
	fmt.Printf("Result: %s\n", result)
}
