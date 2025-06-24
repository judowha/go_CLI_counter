package parser

import (
	"flag"
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
}

func ParseInp() inpArgs {
	var inp inpArgs
	flag.BoolVar(&inp.Case, "case-sensitive", false, "flg for case sensitive")
	flag.Var(&inp.Files, "files", "file addresses that separate by space")
	flag.Parse()
	if len(inp.Files.Names) == 0 {
		panic("No File Input!!")
	}
	return inp
}
