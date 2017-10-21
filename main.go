package main

import (
	"log"

	"github.com/err0r500/go-architect/domain"
	"github.com/err0r500/go-architect/interfaces/mocked"
)

// on pourra virer les paramètres des méthodes et utiliser des propriétés des structs implémentant
// les interfaces dans un second temps, c'est juste pour y voir plus claire au début ...
type TreeExplorer interface {
	GetDirsInTree(rootPath string) (dirPathes *[]string, err error)
	GetFilesInDir(dirPath string) (pathes *[]string, err error)
}

type FileManager interface {
	GetFileContent(path string) (*string, error)
	WriteToFile() error
}

type AstManager interface {
	GetImports(fileContent string) (importsPaths *[]string, err error)
}

type JSONwriter interface {
	ToJSON() (*[]byte, error)
}

type ImportsFinderInteractor struct {
	tE   TreeExplorer
	fM   FileManager
	astM AstManager
}

func main() {
	dummy := ImportsFinderInteractor{
		tE:   mocked.TreeExplorer{},
		fM:   mocked.FileManager{},
		astM: mocked.AstManager{},
	}

	dummy.GetAllImports()
}

// juste un gros bloc pour montrer ma idée initiale, surement naîve
func (i ImportsFinderInteractor) GetAllImports() *[]domain.Pack {
	dirs, _ := i.tE.GetDirsInTree(".")

	for _, dir := range *dirs {
		files, _ := i.tE.GetFilesInDir(dir)

		for _, file := range *files {
			fileContent, _ := i.fM.GetFileContent(file)
			imports, _ := i.astM.GetImports(*fileContent)
			log.Print(imports)
		}
	}
	return nil
}
