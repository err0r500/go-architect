package astManager

import (
	"go/parser"
	"go/token"
	"strings"
)

type AstManager struct{}

func (astM AstManager) GetImportsFromFile(filePath string) (importsPaths *[]string, err error) {
	f, _ := parser.ParseFile(token.NewFileSet(), filePath, nil, parser.ImportsOnly)
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
