package twentythree

import (
	"fmt"
	"strconv"
)

func Day10Problem1(inputs []string) string {
	_, steps := Day10TraverseLoop(inputs)

	return strconv.Itoa(steps)
}

func Day10TraverseLoop(pipes []string) (map[string]int, int) {
	x, y := Day10FindStart(pipes)

	oneX, oneY := Day10NextIndex(x, y, x, y, pipes)
	prevOneX, prevOneY := x, y
	twoX, twoY := Day10NextIndex(x, y, oneX, oneY, pipes)
	prevTwoX, prevTwoY := x, y

	pipeCoords := map[string]int{Day10Key(x, y): 1, Day10Key(oneX, oneY): 1, Day10Key(twoX, twoY): 1}

	for i := 2; i > 0; i++ {
		newOneX, newOneY := Day10NextIndex(oneX, oneY, prevOneX, prevOneY, pipes)
		pipeCoords[Day10Key(newOneX, newOneY)]++
		prevOneX, prevOneY = oneX, oneY
		oneX, oneY = newOneX, newOneY

		newTwoX, newTwoY := Day10NextIndex(twoX, twoY, prevTwoX, prevTwoY, pipes)
		pipeCoords[Day10Key(newTwoX, newTwoY)]++
		prevTwoX, prevTwoY = twoX, twoY
		twoX, twoY = newTwoX, newTwoY

		if pipeCoords[Day10Key(oneX, oneY)] > 1 || pipeCoords[Day10Key(twoX, twoY)] > 1 {
			return pipeCoords, i
		}
	}

	return pipeCoords, 0
}

func Day10Key(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func Day10NextIndex(x, y, prevX, prevY int, pipes []string) (int, int) {
	current := pipes[y][x]

	leftX, leftY, _, leftValid := Day10Left(x, y, prevX, prevY, pipes)
	rightX, rightY, _, rightValid := Day10Right(x, y, prevX, prevY, pipes)
	aboveX, aboveY, _, aboveValid := Day10Above(x, y, prevX, prevY, pipes)
	belowX, belowY, _, belowValid := Day10Below(x, y, prevX, prevY, pipes)

	switch current {
	case 'S':
		if aboveValid {
			return aboveX, aboveY
		}
		if belowValid {
			return belowX, belowY
		}
		if leftValid {
			return leftX, leftY
		}
		if rightValid {
			return rightX, rightY
		}
		fallthrough
	case '|':
		if aboveValid {
			return aboveX, aboveY
		}
		if belowValid {
			return belowX, belowY
		}
		fallthrough
	case '-':
		if leftValid {
			return leftX, leftY
		}
		if rightValid {
			return rightX, rightY
		}
		fallthrough
	case '7':
		if belowValid {
			return belowX, belowY
		}
		if leftValid {
			return leftX, leftY
		}
		fallthrough
	case 'J':
		if aboveValid {
			return aboveX, aboveY
		}
		if leftValid {
			return leftX, leftY
		}
		fallthrough
	case 'F':
		if belowValid {
			return belowX, belowY
		}
		if rightValid {
			return rightX, rightY
		}
		fallthrough
	case 'L':
		if aboveValid {
			return aboveX, aboveY
		}
		if rightValid {
			return rightX, rightY
		}
		fallthrough
	case '.':
		fallthrough
	default:
		return x, y
	}
}

func Day10Left(x, y, prevX, prevY int, pipes []string) (leftX, leftY int, left byte, isValid bool) {
	leftX, leftY = x-1, y

	left = pipes[leftY][leftX]

	isValid = leftX >= 0 &&
		!(leftX == prevX && leftY == prevY) &&
		(left == '-' || left == 'F' || left == 'L')

	return
}

func Day10Right(x, y, prevX, prevY int, pipes []string) (rightX, rightY int, right byte, isValid bool) {
	rightX, rightY = x+1, y

	right = pipes[rightY][rightX]

	isValid = rightX < len(pipes[y]) &&
		!(rightX == prevX && rightY == prevY) &&
		(right == '-' || right == 'J' || right == '7')

	return
}

func Day10Above(x, y, prevX, prevY int, pipes []string) (aboveX, aboveY int, above byte, isValid bool) {
	aboveX, aboveY = x, y-1

	above = pipes[aboveY][aboveX]

	isValid = aboveY >= 0 &&
		!(aboveX == prevX && aboveY == prevY) &&
		(above == '|' || above == '7' || above == 'F')

	return
}

func Day10Below(x, y, prevX, prevY int, pipes []string) (belowX, belowY int, below byte, isValid bool) {
	belowX, belowY = x, y+1

	below = pipes[belowY][belowX]

	isValid = belowY < len(pipes) &&
		!(belowX == prevX && belowY == prevY) &&
		(below == '|' || below == 'L' || below == 'J')

	return
}

func Day10FindStart(inputs []string) (int, int) {
	for y, input := range inputs {
		for x, ch := range input {
			if ch == 'S' {
				return x, y
			}
		}
	}
	return 0, 0
}
