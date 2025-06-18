package parser

import (
	"flag"
	"strings"
)

type fileNameList struct {
	names []string
}

func (fList *fileNameList) String() string {
	return strings.Join(fList.names, " ")
}

func (fList *fileNameList) Set(s string) error {
	fList.names = strings.Split(s, " ")
	return nil
}

type inpArgs struct {
	files         fileNameList
	caseSensitive bool
}

func ParseInp() inpArgs {
	var inp inpArgs
	inp.caseSensitive = *flag.Bool("case-sensitive", false, "flg for case sensitive")
	flag.Var(&inp.files, "files", "file addresses that separate by space")
	flag.Parse()
	if len(inp.files.names) == 0 {
		panic("No File Input!!")
	}
	return inp
}
