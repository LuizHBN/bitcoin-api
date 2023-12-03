package utils

import (
	"strconv"
)

func ParseStringToInt(value string) int64 {
	parsedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	return parsedValue
}

func IntToString(value int) string {
	parsedValue := strconv.FormatInt(int64(value), 10)

	return parsedValue
}
