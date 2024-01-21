package common

import (
	"fmt"

	c22 "github.com/glennhartmann/aoclib/common"
	"golang.org/x/exp/constraints"
)

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func DigitToInt(b byte) int {
	return int(b - '0')
}

func Panicf(fmtStr string, args ...any) {
	panic(fmt.Sprintf(fmtStr, args...))
}

func AddSentinal(lines []string, c string) []string {
	line := c22.Padding(c, len(lines[0]))
	lines = append([]string{line}, lines...)
	lines = append(lines, line)

	for i := range lines {
		lines[i] = fmt.Sprintf("%s%s%s", c, lines[i], c)
	}

	return lines
}

func AddSentinal2[T any](lines [][]T, c T) [][]T {
	line := make([]T, len(lines[0]))
	for i := range line {
		line[i] = c
	}
	lines = append([][]T{line}, lines...)
	lines = append(lines, line)

	for i := range lines {
		lines[i] = append([]T{c}, lines[i]...)
		lines[i] = append(lines[i], c)
	}

	return lines
}

type Equatable interface {
	constraints.Complex | constraints.Float | constraints.Integer | ~string | ~bool | ~byte
}

func SplitSlice[T Equatable](slice []T, sep []T) [][]T {
	ret := make([][]T, 0, len(slice)/5)
	for {
		i := SliceIndex(slice, sep)
		if i == -1 {
			ret = append(ret, slice)
			break
		}

		ret = append(ret, slice[:i])
		slice = slice[i+1:]
	}
	return ret
}

func SliceIndex[T Equatable](slice []T, target []T) int {
outer:
	for i := 0; i < len(slice)-len(target)+1; i++ {
		for j := range target {
			if slice[i+j] != target[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

func StringSliceToByteSlice2(strs []string) [][]byte {
	ret := make([][]byte, len(strs))
	for i := range ret {
		ret[i] = []byte(strs[i])
	}
	return ret
}

func ByteSlice2ToStringSlice(bs [][]byte) []string {
	ret := make([]string, len(bs))
	for i := range ret {
		ret[i] = string(bs[i])
	}
	return ret
}
