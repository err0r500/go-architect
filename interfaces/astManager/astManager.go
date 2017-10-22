package AstManager

import (
	"go/parser"
	"go/token"
	"strings"
)

type AstManager struct{}

func (astM AstManager) GetImports(fileContent string) (importsPaths *[]string, err error) {
	// could be used directly with file :
	// set the real path instead of "dummyPath"
	// and nil instead of the fileContent
	f, _ := parser.ParseFile(token.NewFileSet(), "dummyPath", fileContent, parser.ImportsOnly)
	if f == nil {
		return nil, nil
	}

	ss := []string{}
	for _, importSpec := range f.Imports {
		// also removes the quotes from the returned string
		ss = append(ss, strings.Replace(importSpec.Path.Value, "\"", "", -1))
	}

	return &ss, nil
}
