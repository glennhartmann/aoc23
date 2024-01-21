package common_05

import (
	"log"
	"strings"

	"github.com/glennhartmann/aoclib/must"
)

type RangeMap struct {
	Dst  int64
	Src  int64
	RLen int64
}

func ParseInput() ([]int64, [][]RangeMap) {
	lines := must.GetFullInput()

	seeds := must.ParseListOfNumbers64(strings.Split(lines[0], ": ")[1], " ")
	log.Printf("seeds: %v", seeds)

	maps := [][]RangeMap{
		make([]RangeMap, 0, 50),
		make([]RangeMap, 0, 50),
		make([]RangeMap, 0, 50),
		make([]RangeMap, 0, 50),
		make([]RangeMap, 0, 50),
		make([]RangeMap, 0, 50),
		make([]RangeMap, 0, 50),
	}

	mapIndex := 0
	i := 3
	for i < len(lines) {
		if lines[i] == "" {
			mapIndex++
			i += 2
			continue
		}

		rSpl := must.ParseListOfNumbers64(lines[i], " ")
		rMap := RangeMap{
			Dst:  rSpl[0],
			Src:  rSpl[1],
			RLen: rSpl[2],
		}

		maps[mapIndex] = append(maps[mapIndex], rMap)
		log.Printf("adding RangeMap[%d]: %v", mapIndex, rMap)

		i++
	}

	return seeds, maps
}
