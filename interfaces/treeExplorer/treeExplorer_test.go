package treeExplorer

import (
	"testing"

	"github.com/err0r500/go-architect/testHelpers"
)

func TestGetDirsInTree(t *testing.T) {
	// should be relative to $GOPATH ? -> build.Default.GOPATH
	expected := []string{
		"../../testHelpers/test_data",
		"../../testHelpers/test_data/folder1",
	}

	tE := TreeExplorer{}
	returned, err := tE.GetDirsInTree("../../testHelpers/test_data")
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
	returned, err := tE.GetFilesInDir("../../testHelpers/test_data")
	if err != nil {
		t.Error(err)
	}

	if err := testHelpers.CheckStringSliceEqual(*returned, expected); err != nil {
		t.Error(err)
	}
}
