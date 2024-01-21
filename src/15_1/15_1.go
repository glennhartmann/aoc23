package main

import (
	"log"
	"strings"

	"github.com/glennhartmann/aoclib/must"
)

func main() {
	lines := must.GetFullInput()
	sp := strings.Split(lines[0], ",")
	sum := 0
	for i, step := range sp {
		h := hash(step)
		log.Printf("step %d (%q) hash = %d", i, step, h)
		sum += h
	}
	log.Printf("sum: %d", sum)
}

func hash(s string) int {
	r := 0
	for i := 0; i < len(s); i++ {
		r = (r + int(s[i])) * 17 % 256
	}
	return r
}
