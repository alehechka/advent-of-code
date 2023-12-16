package twentythree_test

import (
	twentythree "advent/solutions/2023"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day16input = []string{
	`.|...\....`,
	`|.-.\.....`,
	`.....|-...`,
	`........|.`,
	`..........`,
	`.........\`,
	`..../.\\..`,
	`.-.-/..|..`,
	`.|....-|.\`,
	`..//.|....`,
}

func Test_Day16Problem1(t *testing.T) {
	assert.Equal(t, "46", twentythree.Day16Problem1(day16input))
}
