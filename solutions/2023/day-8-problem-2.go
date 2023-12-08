package twentythree

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func Day8Problem2(inputs []string) string {
	directions := inputs[0]
	nodes := Day8ParseNodes(inputs[2:])

	distances := make([]int, 0)
	for _, node := range nodes {
		if node.Base[2] == 'A' {
			distances = append(distances, nodes.TraverseToSuffix(directions, node.Base, 'Z'))
		}
	}

	// return fmt.Sprintf("%v", distances)
	return strconv.Itoa(utils.LeastCommonMultipleMany(distances...))
}

func (m Day8NodeMap) TraverseToSuffix(directions string, start string, suffix byte) (steps int) {
	current := start
	var iterations int
	for {
		for _, ch := range directions {
			node, ok := m[current]
			if !ok {
				return
			}
			if node.Base[2] == suffix {
				return
			}
			steps++
			if ch == 'L' {
				current = node.Left
			} else {
				current = node.Right
			}
		}
		iterations++
		// fmt.Println("Exhausted directions, going again: ", iterations)
	}
}

func Day8Problem2_FullTraversal(inputs []string) string {
	directions := inputs[0]
	nodes := Day8ParseNodes(inputs[2:])

	starts := make(map[string]string)
	for _, node := range nodes {
		if node.Base[2] == 'A' {
			starts[node.Base] = node.Base
		}
	}

	steps := nodes.TraverseMany(directions, starts)

	return strconv.Itoa(steps)
}

func (m Day8NodeMap) TraverseMany(directions string, starts map[string]string) (steps int) {
	var iterations int
	for {
		for _, ch := range directions {
			steps++

			var endsWithZ int
			for key, current := range starts {
				node := m[current]
				var next string
				if ch == 'L' {
					next = node.Left
				} else {
					next = node.Right
				}
				starts[key] = next
				if next[2] == 'Z' {
					endsWithZ++
				}
			}

			if endsWithZ >= 4 {
				fmt.Println("Ends with zero: ", endsWithZ)
			}
			if endsWithZ == len(starts) {
				return
			}
		}
		iterations++
		fmt.Println("Exhausted directions, going again: ", iterations, steps)
	}
}
