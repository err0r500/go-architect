package analyzer

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
)

func AnnalyzeFile() {
	// src is the input for which we want to inspect the AST.
	src := `
	package thePackageName
	import (
		"flag"
		"fmt"
		"path/filepath"
	
		"github.com/err0r500/codeAnalyzer/analyzer"
	)
	const c = 1.0
	var X = f(3.14)*2 + c
	func myfunc(myInterface){
		return
	}
	type myInterface interface {
		doThis()
	}
	`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}
	ss := []string{}

	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
		var s string

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
		if s != "" {
			fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		fmt.Println("---")
		return true
	})
	log.Print(ss)
}
