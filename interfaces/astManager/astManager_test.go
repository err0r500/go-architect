package astManager

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// commenting just to test the next one with a verbose output

// var src = `
// package thePackageName
// import (
// 	"flag"
// 	"fmt"
// 	"path/filepath"

// 	mA "github.com/err0r500/codeAnalyzer/analyzer"
// )
// const c = 1.0
// var X = f(3.14)*2 + c
// func myfunc(myInterface){
// 	return
// }
// type myInterface interface {
// 	doThis()
// }
// `

// func TestAstManager_GetImports(t *testing.T) {
// 	expected := []string{
// 		"flag",
// 		"path/filepath",
// 		"fmt",
// 		"github.com/err0r500/codeAnalyzer/analyzer",
// 	}

// 	testFilePath := "./testFile.go"

// 	ioutil.WriteFile(testFilePath, []byte(src), 0644)
// 	defer os.Remove(testFilePath)
// 	astM := AstManager{}
// 	returned, err := astM.GetImportsFromFile(testFilePath)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if err := testHelpers.CheckStringSliceEqual(expected, *returned); err != nil {
// 		t.Error(err)
// 	}
// }

func TestDe(t *testing.T) {
	src := `
		package p
		import (
			"io"
			log "github.com/sirupsen/logrus"
		)
		type customStruc struct {}

		const c = 1.0
		var X = f(3.14)*2 + c

		type myInterface interface {
			doThis() error
			doThat(what bool) (how int, nope customStruc)
			andThat(some io.Reader) (anImport log.Fields)
	 	}
	`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "anyname", src, 0)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		inspect(n)
		// switch x := n.(type) {
		// case *ast.InterfaceType:
		// 	fmt.Print("\nInterface contains : \n") // pas grand chose d'utile dans ast.InterfaceType, on dirait...

		// case *ast.FuncDecl:
		// 	fmt.Sprintf("FUNC => %#v", x.Type.Params)
		// 	// case *ast.Ident:
		// 	// 	fmt.Sprintf("FUNC => %#v", x.Type.Params)
		// 	// 	s = x.Name
		// }

		return true
	})

}
