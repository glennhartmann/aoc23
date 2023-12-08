// Package must contains helper functions that panic if they have errors.
package must

import (
	"strconv"

	"github.com/glennhartmann/aoc23/src/common"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		common.Panicf("invalid int for Atoi: %s", s)
	}
	return i
}

func Atoi64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		common.Panicf("invalid int64 for Atoi64: %s", s)
	}
	return i
}
