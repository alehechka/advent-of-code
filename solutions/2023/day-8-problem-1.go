package twentythree

import (
	"fmt"
	"strconv"
)

func Day8Problem1(inputs []string) string {
	directions := inputs[0]
	nodes := Day8ParseNodes(inputs[2:])

	steps := nodes.Traverse(directions, "AAA", "ZZZ")

	return strconv.Itoa(steps)
}

type Day8NodeMap map[string]Day8Node

func (m Day8NodeMap) Traverse(directions string, start string, end string) (steps int) {
	current := start
	var iterations int
	for {
		for _, ch := range directions {
			node, ok := m[current]
			if !ok {
				return
			}
			if node.Base == end {
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
		fmt.Println("Exhausted directions, going again: ", iterations)
	}
}

type Day8Node struct {
	Base  string
	Left  string
	Right string
}

func Day8ParseNodes(inputs []string) Day8NodeMap {
	nodes := make(Day8NodeMap)

	for _, input := range inputs {
		node := Day8ParseNode(input)
		nodes[node.Base] = node
	}

	return nodes
}

func Day8ParseNode(input string) (node Day8Node) {
	node.Base = input[0:3]
	node.Left = input[7:10]
	node.Right = input[12:15]
	return
}
