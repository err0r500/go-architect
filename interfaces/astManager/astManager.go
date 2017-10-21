package AstManager

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
)

type AstManager struct{}

func (astM AstManager) GetImports(fileContent string) (importsPaths *[]string, err error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", fileContent, 0)
	if err != nil {
		return nil, err
	}
	ss := []string{}

	ast.Inspect(f, func(n ast.Node) bool {
		s := ""

		fmt.Println(reflect.TypeOf(n))

		switch x := n.(type) {
		case *ast.Ident:
			s = "curPackageName : " + x.Name
			ss = append(ss, s)
			// case *ast.ImportSpec:
			// 	s = "import : " + x.Path.Value
			// 	break
			// case *ast.BasicLit:
			// 	s = "literal : " + x.Value + " -> " + x.Kind.String()
			// case *ast.CallExpr:
			// 	s = x.Args x.Name + " -> " + fmt.Sprintf("%s", x.Obj)
		}
		return true
	})

	log.Print(ss)
	return &ss, nil
}
