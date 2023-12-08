package main

import (
	"log"

	c05 "github.com/glennhartmann/aoc23/src/common_05"
)

type normalRange struct {
	start int64
	rLen  int64
}

func main() {
	seeds, maps := c05.ParseInput()

	seedRanges := make([]normalRange, 0, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, normalRange{
			start: seeds[i],
			rLen:  seeds[i+1],
		})
		log.Printf("adding normalRange of seeds: %v", seedRanges[len(seedRanges)-1])
	}

	mappedVal := seedRanges
	for _, m := range maps {
		mappedVal = resolveMaps(mappedVal, m)
	}

	log.Printf("seed ranges %v correspond to location ranges %v", seedRanges, mappedVal)

	minLocation := int64(-1)
	for _, lcn := range mappedVal {
		if minLocation == int64(-1) || lcn.start < minLocation {
			minLocation = lcn.start
		}
	}

	log.Printf("min location: %d", minLocation)
}

var badRange = normalRange{-1, -1}

func resolveMaps(v []normalRange, maps []c05.RangeMap) []normalRange {
	ret := make([]normalRange, 0, 500)
outer:
	for i := 0; i < len(v); i++ {
		rng := v[i]
		for _, rMap := range maps {
			left, match, right := findInRangeMap(rng, rMap)
			if match == badRange {
				continue
			}

			matchCpy := match
			match.start = rMap.Dst + match.start - rMap.Src
			log.Printf("converting match from src space (%v) to dst space (%v)", matchCpy, match)

			ret = append(ret, match)
			if left != badRange {
				v = append(v, left)
			}
			if right != badRange {
				v = append(v, right)
			}
			continue outer
		}
		ret = append(ret, rng)
	}
	return ret
}

func findInRangeMap(v normalRange, m c05.RangeMap) (left, match, right normalRange) {
	match = getOverlap(v, m)
	if match == badRange {
		return badRange, badRange, badRange
	}

	left, right = badRange, badRange

	if match.start > v.start {
		left = normalRange{
			start: v.start,
			rLen:  match.start - v.start,
		}
	}

	matchEnd := match.start + match.rLen - 1
	vEnd := v.start + v.rLen - 1
	if matchEnd < vEnd {
		right = normalRange{
			start: matchEnd + 1,
			rLen:  vEnd - matchEnd,
		}
	}

	log.Printf("splitting {%v; %v} into {%v; %v; %v}", v, m, left, match, right)
	return left, match, right
}

func getOverlap(v normalRange, m c05.RangeMap) normalRange {
	mEnd := m.Src + m.RLen - 1
	vEnd := v.start + v.rLen - 1
	if v.start > mEnd || vEnd < m.Src {
		return badRange
	}

	start := max(v.start, m.Src)
	return normalRange{
		start: start,
		rLen:  min(vEnd, mEnd) - start + 1,
	}
}
