package utils

import (
	"fmt"
	"testing"

	"github.com/judowha/go_CLI_counter.git/reader"
)

func TestCleanText(t *testing.T) {
	result := CleanText(reader.ReadFiles([]string{"../test_data/test1.txt", "../test_data/test2.txt"}))
	fmt.Println(result)
}
