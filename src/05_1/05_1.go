package main

import (
	"log"

	c05 "github.com/glennhartmann/aoc23/src/common_05"
)

func main() {
	seeds, maps := c05.ParseInput()

	minLocation := int64(-1)
	for _, seed := range seeds {
		mappedVal := seed
		for _, m := range maps {
			mappedVal = resolveMaps(mappedVal, m)
		}
		log.Printf("seed %d corresponds to location %d", seed, mappedVal)

		if minLocation == int64(-1) || mappedVal < minLocation {
			minLocation = mappedVal
		}

		log.Printf("min location so far: %d", minLocation)
	}

	log.Printf("min location: %d", minLocation)
}

func resolveMaps(v int64, maps []c05.RangeMap) int64 {
	for _, rMap := range maps {
		if i := findInRangeMap(v, rMap); i != -1 {
			return i
		}
	}
	return v
}

func findInRangeMap(v int64, m c05.RangeMap) int64 {
	if v >= m.Src && v < m.Src+m.RLen {
		return m.Dst + v - m.Src
	}
	return -1
}
