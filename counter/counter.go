package counter

import (
	"strings"
)

func CountFreq(s string, caseSens bool) (result map[string]int) {
	result = make(map[string]int)
	if !caseSens {
		s = strings.ToLower(s)
	}
	strList := strings.Split(s, " ")
	for _, str := range strList {
		result[str] += 1
	}
	return result
}
