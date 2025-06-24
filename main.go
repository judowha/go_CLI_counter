package main

import (
	"fmt"

	"github.com/judowha/go_CLI_counter.git/counter"
	"github.com/judowha/go_CLI_counter.git/parser"
	"github.com/judowha/go_CLI_counter.git/reader"
)

func main() {
	inp := parser.ParseInp()
	fmt.Println(inp)
	contain := reader.ReadFiles(inp.Files.Names)
	result := counter.CountFreq(contain, inp.Case)
	fmt.Println(result)
}
