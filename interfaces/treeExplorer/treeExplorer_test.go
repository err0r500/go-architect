package mocked

import (
	"testing"

	"github.com/err0r500/go-architect/testHelpers"
)

func TestRelativePath(t *testing.T) {
	// FixMe, should return path from $GOPATH ?
	expected := []string{
		"./test_data",
		"test_data/folder1",
	}

	tE := TreeExplorer{}
	returned, err := tE.GetDirsInTree("./test_data")
	if err != nil {
		t.Error(err)
	}

	if err := testHelpers.CheckStringSliceEqual(*returned, expected); err != nil {
		t.Error(err)
	}
}
