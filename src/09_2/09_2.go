package main

import (
	"log"

	"github.com/glennhartmann/aoclib/must"
)

func main() {
	lines := must.GetFullInput()

	sum := int64(0)
	for i, line := range lines {
		nums := must.ParseListOfNumbers64(line, " ")
		nums = append([]int64{0}, nums...)
		log.Printf("line %d: %v", i, nums)

		diffs := make([][]int64, 0, 5)
		diffs = append(diffs, nums)
		for {
			diffIndex := len(diffs) - 1

			diff := make([]int64, 0, len(diffs[diffIndex]))
			diff = append(diff, 0)

			for i := 2; i < len(diffs[diffIndex]); i++ {
				diff = append(diff, diffs[diffIndex][i]-diffs[diffIndex][i-1])
			}

			diffs = append(diffs, diff)
			log.Printf("line %d: diffs[%d] = %v", i, diffIndex+1, diff)

			if isAllZeroes(diff) {
				break
			}
		}

		diffs[len(diffs)-1] = append(diffs[len(diffs)-1], 0)
		for j := len(diffs) - 2; j >= 0; j-- {
			diffs[j][0] = diffs[j][1] - diffs[j+1][0]
			log.Printf("line %d: expanded diffs[%d] = %v", i, j, diffs[j])
		}

		sum += diffs[0][0]
	}

	log.Printf("final sum: %d", sum)
}

func isAllZeroes(nums []int64) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}
