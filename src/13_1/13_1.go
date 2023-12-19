package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/glennhartmann/aoc23/src/common"
	"github.com/glennhartmann/aoc23/src/common/must"
)

func main() {
	lines := must.GetFullInput()
	spLines := common.SplitSlice(lines, []string{""})

	sum := 0
	for i, pattern := range spLines {
		log.Printf("pattern %d:\n%s", i, strings.Join(pattern, "\n"))

		sym := findSymmetry(pattern)
		if sym == -1 {
			sym = findSymmetry(cols(pattern))
			sum += sym + 1
			log.Printf("horizontal symmetry found at %d", sym)
		} else {
			sum += 100 * (sym + 1)
			log.Printf("vertical symmetry found at %d", sym)
		}
	}
	log.Printf("sum: %d", sum)
}

// returns the index of the row/col before (ie, to the left of or
// above) the line of symmetry.
func findSymmetry(rc []string) int {
	startSym := findSymmetryFromStart(rc)
	if startSym == -1 {
		cp := slices.Clone(rc)
		slices.Reverse(cp)
		endSym := findSymmetryFromStart(cp)
		if endSym == -1 {
			return -1
		}
		return len(cp) - endSym - 2
	}
	return startSym
}

func findSymmetryFromStart(rc []string) int {
	start := rc[0]
	index := 1
	for {
		i := slices.Index(rc[index:], start)
		if i == -1 {
			break
		}
		if isPalindrome(rc[:i+index+1]) {
			return (i+index+1)/2 - 1
		}
		index += i + 1
	}
	return -1
}

func isPalindrome(s []string) bool {
	if len(s)%2 != 0 {
		// real palindromes can have odd lengths, but for this question
		// we need a mirror between 2 rows/cols
		return false
	}
	for i := 0; i < len(s)/2; i++ {
		left := s[i]
		right := s[len(s)-i-1]
		if left != right {
			return false
		}
	}
	return true
}

func cols(rows []string) []string {
	sbs := make([]strings.Builder, len(rows[0]))
	for row := range rows {
		for col := 0; col < len(rows[0]); col++ {
			fmt.Fprintf(&sbs[col], "%c", rows[row][col])
		}
	}
	ret := make([]string, len(sbs))
	for i := range ret {
		ret[i] = sbs[i].String()
	}
	return ret
}
