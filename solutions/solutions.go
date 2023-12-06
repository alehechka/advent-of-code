package solutions

import (
	twentythree "advent/solutions/2023"
)

var Solutions = map[int]SolutionArray{
	2023: Solutions2023,
}

type Solution func([]string) string

type SolutionArray [][]Solution

var Solutions2023 = SolutionArray{
	{twentythree.Day1Problem1, twentythree.Day1Problem2},
	{twentythree.Day2Problem1, twentythree.Day2Problem2},
	{twentythree.Day3Problem1, twentythree.Day3Problem2},
	{twentythree.Day4Problem1, twentythree.Day4Problem2},
	{twentythree.Day5Problem1, twentythree.Day5Problem2},
	{twentythree.Day6Problem1, twentythree.Day6Problem2},
}
