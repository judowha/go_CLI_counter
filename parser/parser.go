package parser

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type fileNameList struct {
	Names []string
}

func (fList *fileNameList) String() string {
	return strings.Join(fList.Names, " ")
}

func (fList *fileNameList) Set(s string) error {
	fList.Names = strings.Split(s, " ")
	return nil
}

type inpArgs struct {
	Files fileNameList
	Case  bool
	Sort  string
}

var allowSort = map[string]bool{
	"freq":  true,
	"alpha": true,
}

func ParseInp() (inp inpArgs) {
	flag.StringVar(&inp.Sort, "sort", "freq", "critira for sorting, possible options: freq|aplha")
	flag.BoolVar(&inp.Case, "case-sensitive", false, "flg for case sensitive")
	flag.Var(&inp.Files, "files", "file addresses that separate by space")
	flag.Parse()

	if !allowSort[inp.Sort] {
		fmt.Fprintf(os.Stderr, "Invalid Sort: %s. Choose from freq | aplha.\n", inp.Sort)
		os.Exit(1)
	}

	if len(inp.Files.Names) == 0 {
		fmt.Fprintf(os.Stderr, "No file input\n")
		os.Exit(1)
	}
	return inp
}
