package AstManager

import (
	"testing"

	"github.com/err0r500/go-architect/testHelpers"
)

func TestAstManager_GetImports(t *testing.T) {
	expected := []string{
		"flag",
		"fmt",
		"path/filepath",
		"github.com/err0r500/codeAnalyzer/analyzer",
	}

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
	astM := AstManager{}

	returned, err := astM.GetImports(src)
	if err != nil {
		t.Error(err)
	}

	if err := testHelpers.CheckStringSliceEqual(expected, *returned); err != nil {
		// t.Error(err)
	}
}
