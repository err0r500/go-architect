package treeExplorer

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type TreeExplorer struct{}

// GetDirsInTree : very simply tested : just returns relative pathes ...
func (fE TreeExplorer) GetDirsInTree(rootPath string) (*[]string, error) {
	dirs := []string{}

	visit := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			dirs = append(dirs, filepath.Clean(path))
		}
		return nil
	}

	err := filepath.Walk(rootPath, visit)

	return &dirs, err
}

func (fE TreeExplorer) GetFilesInDir(dirPath string) (pathes *[]string, err error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	ff := []string{}
	for _, f := range files {
		if isInterestingFile(f) {
			ff = append(ff, f.Name())
		}
	}

	return &ff, nil
}

func isInterestingFile(f os.FileInfo) bool {
	if f.IsDir() {
		return false
	}
	if filepath.Ext(f.Name()) != ".go" {
		return false
	}
	return true
}
