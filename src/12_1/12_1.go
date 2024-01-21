package main

import (
	"log"
	"strings"

	"github.com/glennhartmann/aoclib/must"
)

func main() {
	lines := must.GetFullInput()

	count := 0
	for i, line := range lines {
		spLine := strings.Split(line, " ")
		pattern := spLine[0]
		groups := must.ParseListOfNumbers(spLine[1], ",")

		intervals := getUsableIntervals(pattern)
		log.Printf("pattern %d intervals: %+v", i, intervals)

		perm := permuteRecursive(pattern, intervals, groups, groups)
		log.Printf("permutations for line %d: %d", i, perm)
		count += perm
	}
	log.Printf("total: %d", count)
}

type interval struct {
	start, length int
}

func getUsableIntervals(p string) []interval {
	intervals := make([]interval, 0, 50)

	intervalStart := 0
	for i := 0; i < len(p); i++ {
		if p[i] == '.' {
			intervals = maybeAddInterval(intervals, intervalStart, i-intervalStart)
			intervalStart = i + 1
		}
	}
	intervals = maybeAddInterval(intervals, intervalStart, len(p)-intervalStart)

	return intervals
}

func maybeAddInterval(intervals []interval, start, length int) []interval {
	if length > 0 {
		intervals = append(intervals, interval{
			start:  start,
			length: length,
		})
	}
	return intervals
}

func permuteRecursive(p string, intervals []interval, groups, originalGroups []int) int {
	if len(groups) == 0 {
		newP := format(p)
		if !doubleCheck(newP, originalGroups) {
			log.Printf("found BAD permutation: %s", newP)
			return 0
		}
		log.Printf("found good permutation: %s", newP)
		return 1
	}

	permutations := 0
	for i := 0; i < len(p); i++ {
		if fits(p, i, intervals, groups[0]) {
			permutations += permuteRecursive(updatePattern(p, i, groups[0]), splitIntervals(intervals, i, groups[0]), groups[1:], originalGroups)
		}
	}
	return permutations
}

func doubleCheck(p string, groups []int) bool {
	intervals := getUsableIntervals(p)
	log.Printf("doublecheck intervals %+v", intervals)
	if len(intervals) != len(groups) {
		return false
	}

	for i := range intervals {
		if intervals[i].length != groups[i] {
			return false
		}
	}

	return true
}

func fits(p string, i int, intervals []interval, group int) bool {
	rInterval, _ := getRelevantInterval(i, intervals)
	if rInterval.length == -1 {
		return false
	}

	if i+group-1 > rInterval.start+rInterval.length-1 {
		return false
	}

	if (i > 0 && p[i-1] == '#') || (i+group < len(p)-1 && p[i+group] == '#') {
		return false
	}

	if skippedBroken(p, i, rInterval) {
		return false
	}

	return true
}

func skippedBroken(p string, i int, rInterval interval) bool {
	for c := rInterval.start; c < i; c++ {
		if p[c] == '#' {
			return true
		}
	}
	return false
}

func getRelevantInterval(i int, intervals []interval) (interval, int) {
	for intervalIndex, rInterval := range intervals {
		if i >= rInterval.start && i <= rInterval.start+rInterval.length-1 {
			return rInterval, intervalIndex
		}
	}
	return interval{start: -1, length: -1}, -1
}

func updatePattern(p string, i, group int) string {
	pb := []byte(p)

	for c := i; c < i+group; c++ {
		pb[c] = '#'
	}

	return string(pb)
}

func splitIntervals(intervals []interval, i, group int) []interval {
	rInterval, intervalIndex := getRelevantInterval(i, intervals)

	nIntervals := make([]interval, 0, len(intervals)+1)
	nIntervals = append(nIntervals, splitInterval(rInterval, i, group)...)

	for i := intervalIndex + 1; i < len(intervals); i++ {
		nIntervals = append(nIntervals, intervals[i])
	}

	return nIntervals
}

func splitInterval(rInterval interval, i, group int) []interval {
	intervals := make([]interval, 0, 1)

	if i+group-1 < rInterval.start+rInterval.length-1 {
		intervals = append(intervals, interval{
			start:  i + group,
			length: rInterval.start + rInterval.length - (i + group),
		})
	}

	return intervals
}

func format(p string) string {
	return strings.ReplaceAll(p, "?", ".")
}
