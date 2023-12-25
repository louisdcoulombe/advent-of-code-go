package util

import (
	"strings"

	"github.com/louisdcoulombe/advent-of-code-go/cast"
)

func ReplaceAtIndex(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func StringsToInts(str string, sep string) (result []int) {
	parts := strings.Split(str, sep)
	for _, s := range parts {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		result = append(result, cast.ToInt(s))
	}
	return result
}
