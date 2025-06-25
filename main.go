package main

import (
	"fmt"

	"github.com/judowha/go_CLI_counter.git/counter"
	"github.com/judowha/go_CLI_counter.git/parser"
	"github.com/judowha/go_CLI_counter.git/reader"
	"github.com/judowha/go_CLI_counter.git/utils"
)

func main() {
	inp := parser.ParseInp()
	contain := reader.ReadFiles(inp.Files.Names)
	contain = utils.CleanText(contain)
	result := counter.CountFreq(contain, inp.Case)
	counter.SortResult(result, inp.Sort)
	fmt.Println(result)
}
