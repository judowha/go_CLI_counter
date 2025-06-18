package reader

import (
	"fmt"
	"os"
	"strings"
)

func ReadFiles(files []string) string {
	var sb strings.Builder
	for _, file := range files {
		contain, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Fail to open file" + file)
			continue
		}
		sb.WriteString(string(contain))
		sb.WriteString(" ")
	}
	return sb.String()
}
