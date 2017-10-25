package astManager

import (
	"fmt"
	"go/ast"
)

var sep = "\n"

func printSep() {
	sep += "-"
	fmt.Print(sep)
}

func inspect(n ast.Node) {
	switch n.(type) {
	case *ast.InterfaceType:
		inspectInterface(n.(*ast.InterfaceType))
	case *ast.FuncType:
		inspectFuncType(n.(*ast.FuncType))
	case *ast.Ident:
		inspectIdentType(n.(*ast.Ident))
	case *ast.SelectorExpr:
		inspectSelectorExpr(n.(*ast.SelectorExpr))
	case *ast.Field:
		inspectField(n.(*ast.Field))
	case *ast.FuncDecl:
		inspectFuncDecl(n.(*ast.FuncDecl))
	}
}

func inspectInterface(i *ast.InterfaceType) {
	if i == nil {
		return
	}
	fmt.Print("\nInterface contains :\n")
	for _, method := range i.Methods.List {
		printSep()
		fmt.Printf("method => %s\n", method.Names)
		inspect(method.Type)
	}
}

func inspectFuncDecl(f *ast.FuncDecl) {
	if f == nil {
		return
	}
	printSep()

	fmt.Sprintf("FUNC => %#v", f)
}

func inspectFuncType(f *ast.FuncType) {
	if f == nil {
		return
	}

	printSep()
	fmt.Print("params  => \n")
	for _, param := range f.Params.List {
		inspectField(param)
	}

	fmt.Print("returns => \n")
	for _, ret := range f.Results.List {
		inspectField(ret)
	}
}

func inspectField(f *ast.Field) {
	if f == nil {
		return
	}

	fmt.Print("Field => ")
	inspect(f.Type)
}

func inspectIdentType(ident *ast.Ident) {
	if ident == nil {
		return
	}
	fmt.Print("Ident => ")
	fmt.Printf("%#v\n", ident.Name)
}

func inspectSelectorExpr(selector *ast.SelectorExpr) {
	fmt.Print("Selector Expression => ")
	if selector == nil {
		return
	}
	inspect(selector.X)
}
