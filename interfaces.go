package main

import "github.com/err0r500/go-architect/domain"

// on pourra virer les paramètres des méthodes et utiliser des propriétés des structs implémentant
// les interfaces dans un second temps, c'est juste pour y voir plus claire au début ...
type TreeExplorer interface {
	GetDirsInTree(rootPath string) (dirPathes *[]string, err error)
	GetFilesInDir(dirPath string) (pathes *[]string, err error)
}

type FileManager interface {
	GetFileContent(domain.File) (*string, error)
	Write(domain.File) error
}

type AstManager interface {
	GetImports(fileContent string) (importsPaths *[]string, err error)
}

type JSONwriter interface {
	ToJSON(dG *domain.Graph) (string, error)
}
