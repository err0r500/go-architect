package treeExplorer

import (
	"testing"

	"github.com/err0r500/go-architect/testHelpers"
)

func TestGetDirsInTree(t *testing.T) {
	expected := []string{
		"test_data",
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

func TestGetFilesInDir(t *testing.T) {
	expected := []string{
		"f1.go",
		"f_3-Source.go",
	}

	tE := TreeExplorer{}
	returned, err := tE.GetFilesInDir("./test_data")
	if err != nil {
		t.Error(err)
	}

	if err := testHelpers.CheckStringSliceEqual(*returned, expected); err != nil {
		t.Error(err)
	}
}
