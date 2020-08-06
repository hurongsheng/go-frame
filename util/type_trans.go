package util

import "strconv"

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func IntToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}
