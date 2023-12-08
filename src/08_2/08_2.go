package main

import (
	"log"
	"regexp"

	"github.com/glennhartmann/aoc23/src/common/must"
)

var rx = regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)

type graph struct {
	name  string
	left  *graph
	right *graph
}

// running this makes it clear that the 6 start paths hit end paths every
// 19241, 18157, 19783, 16531, 21409, and 14363 steps, respectively.
// The LCM of all those is 24035773251517.
func main() {
	lines := must.GetFullInput()

	path := lines[0]
	log.Printf("path: %s", path)

	graphNodes := make(map[string]*graph, len(lines[2:]))
	startNodes := make([]*graph, 0, len(lines[2:])/4)
	for i, node := range lines[2:] {
		m := must.FindStringSubmatch(rx, node, 4)
		name := m[1]
		leftName := m[2]
		rightName := m[3]

		initializeNodeIfNil(graphNodes, name)
		initializeNodeIfNil(graphNodes, leftName)
		initializeNodeIfNil(graphNodes, rightName)

		graphNodes[name].name = name
		graphNodes[name].left = graphNodes[leftName]
		graphNodes[name].right = graphNodes[rightName]

		log.Printf("node %d: adding %s = (%s, %s)", i, graphNodes[name].name, graphNodes[name].left.name, graphNodes[name].right.name)

		if name[2] == 'A' {
			startNodes = append(startNodes, graphNodes[name])
			log.Printf("%s is a start node", name)
		}
	}

	steps := 0
	pathIndex := 0
	currentNodes := startNodes
	endNodeSteps := make([][]int, len(startNodes))
	endNodeStepDiffs := make([][]int, len(startNodes))
	for {
		if allNodesAreEndNodes(currentNodes) {
			break
		}

		for i, currentNode := range currentNodes {
			if path[pathIndex] == 'L' {
				currentNodes[i] = currentNode.left
			} else {
				currentNodes[i] = currentNode.right
			}

			if currentNodes[i].name[2] == 'Z' {
				endNodeSteps[i] = append(endNodeSteps[i], steps+1)
				if len(endNodeSteps[i]) > 1 {
					endNodeStepDiffs[i] = append(endNodeStepDiffs[i], endNodeSteps[i][len(endNodeSteps[i])-1]-endNodeSteps[i][len(endNodeSteps[i])-2])
				}
			}

			log.Printf("currentNodes[%d] on end nodes at steps %v", i, endNodeSteps[i])
			log.Printf("currentNodes[%d] end node step diffs: %v", i, endNodeStepDiffs[i])
		}

		steps++
		pathIndex = (pathIndex + 1) % len(path)
	}

	log.Printf("path took %d steps", steps)
}

func initializeNodeIfNil(graphNodes map[string]*graph, name string) {
	if graphNodes[name] == nil {
		graphNodes[name] = &graph{name: name}
	}
}

func allNodesAreEndNodes(nodes []*graph) bool {
	for _, node := range nodes {
		if node.name[2] != 'Z' {
			return false
		}
	}
	return true
}
