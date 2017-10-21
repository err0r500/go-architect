package mocked

import (
	"os"
	"path/filepath"
)

type TreeExplorer struct{}

// GetDirsInTree : very simply tested : just returns relative pathes ...
func (fE TreeExplorer) GetDirsInTree(rootPath string) (*[]string, error) {
	dirs := []string{}

	visit := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			dirs = append(dirs, path)
		}
		return nil
	}

	err := filepath.Walk(rootPath, visit)

	return &dirs, err
}

func (fE TreeExplorer) GetFilesInDir(dirPath string) (pathes *[]string, err error) {
	return nil, nil
}
