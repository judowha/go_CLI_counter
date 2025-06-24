package utils

import (
	"regexp"
)

func CleanText(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	clean := re.ReplaceAllString(s, "")
	re = regexp.MustCompile(`[\s]+`)
	clean = re.ReplaceAllString(clean, " ")
	return clean
}
