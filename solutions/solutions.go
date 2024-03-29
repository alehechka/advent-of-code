package solutions

import (
	fifteen "advent/solutions/2015"
	twentythree "advent/solutions/2023"
)

var Solutions = map[int]SolutionArray{
	2023: Solutions2023,
	2015: Solutions2015,
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
	{twentythree.Day7Problem1, twentythree.Day7Problem2},
	{twentythree.Day8Problem1, twentythree.Day8Problem2},
	{twentythree.Day9Problem1, twentythree.Day9Problem2},
	{twentythree.Day10Problem1, twentythree.Day10Problem2},
	{twentythree.Day11Problem1, twentythree.Day11Problem2},
	{twentythree.Day12Problem1, twentythree.Day12Problem2},
	{twentythree.Day13Problem1, twentythree.Day13Problem2},
	{twentythree.Day14Problem1, twentythree.Day14Problem2},
	{twentythree.Day15Problem1, twentythree.Day15Problem2},
	{twentythree.Day16Problem1, twentythree.Day16Problem2},
	{twentythree.Day17Problem1, twentythree.Day17Problem2},
}

var Solutions2015 = SolutionArray{
	{fifteen.Day1Problem1, fifteen.Day1Problem2},
	{fifteen.Day2Problem1, fifteen.Day2Problem2},
	{fifteen.Day3Problem1, fifteen.Day3Problem2},
	{fifteen.Day4Problem1, fifteen.Day4Problem2},
	{fifteen.Day5Problem1, fifteen.Day5Problem2},
}
