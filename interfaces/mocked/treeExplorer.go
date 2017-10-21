package mocked

type TreeExplorer struct{}

func (fE TreeExplorer) GetDirsInTree(rootPath string) (dirPathes *[]string, err error) {
	return nil, nil
}

func (fE TreeExplorer) GetFilesInDir(dirPath string) (pathes *[]string, err error) {
	return nil, nil
}
