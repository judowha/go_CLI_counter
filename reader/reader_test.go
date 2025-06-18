package reader

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	contain := ReadFiles([]string{"../test_data/test1.txt", "../test_data/test2.txt"})
	fmt.Println(contain)
}
