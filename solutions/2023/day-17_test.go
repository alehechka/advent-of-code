package twentythree_test

import (
	twentythree "advent/solutions/2023"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day17example = []string{
	"2413432311323",
	"3215453535623",
	"3255245654254",
	"3446585845452",
	"4546657867536",
	"1438598798454",
	"4457876987766",
	"3637877979653",
	"4654967986887",
	"4564679986453",
	"1224686865563",
	"2546548887735",
	"4322674655533",
}

func Test_Day17Problem1(t *testing.T) {
	assert.Equal(t, "102", twentythree.Day17Problem1(day17example))
}
