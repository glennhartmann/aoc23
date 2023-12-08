package common

import (
	"fmt"
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
