package util

import (
	"strconv"
	"strings"
)

func ParseInts(s string) []int64 {
	splits := strings.Split(s, "\n")
	out := make([]int64, len(splits))

	var err error
	for i, s := range splits {
		out[i], err = strconv.ParseInt(s, 10, 64)
		PanicIfErr(err)
	}

	return out
}

func MustParseInt(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	PanicIfErr(err)

	return n
}