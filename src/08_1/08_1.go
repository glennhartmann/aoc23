package main

import (
	"log"
	"regexp"

	"github.com/glennhartmann/aoclib/must"
)

var rx = regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)

type graph struct {
	name  string
	left  *graph
	right *graph
}

func main() {
	lines := must.GetFullInput()

	path := lines[0]
	log.Printf("path: %s", path)

	graphNodes := make(map[string]*graph, len(lines[2:]))
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
	}

	steps := 0
	pathIndex := 0
	currentNode := graphNodes["AAA"]
	for currentNode.name != "ZZZ" {
		oldCurrentNode := currentNode
		if path[pathIndex] == 'L' {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}

		log.Printf("moved %c from %s to %s", path[pathIndex], oldCurrentNode.name, currentNode.name)

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
