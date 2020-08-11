package util

import (
	"strconv"
	"strings"
)

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
func StrToInts(s string, sep string) []int {
	arrStr := strings.Split(s, sep)
	arr := make([]int, 0, len(arrStr))
	for _, s := range arrStr {
		arr = append(arr, StrToInt(s))
	}
	return arr
}

func IntToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}
