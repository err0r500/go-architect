package main

import (
	"github.com/err0r500/go-architect/domain"
	mockedAstM "github.com/err0r500/go-architect/interfaces/astManager/mocked"
	mockedFM "github.com/err0r500/go-architect/interfaces/fileManager/mocked"
	mockedTE "github.com/err0r500/go-architect/interfaces/treeExplorer/mocked"
)

type ImportsFinderInteractor struct {
	tE   TreeExplorer
	fM   FileManager
	astM AstManager
}

func main() {
	dummy := ImportsFinderInteractor{
		tE:   mockedTE.TreeExplorer{},
		fM:   mockedFM.FileManager{},
		astM: mockedAstM.AstManager{},
	}

	dummy.GetAllImports()
}

// juste un gros bloc pour montrer l'idée initiale, surement naîve
func (i ImportsFinderInteractor) GetAllImports() *[]domain.Pack {
	dirs, _ := i.tE.GetDirsInTree(".")

	packageList := []domain.Pack{}

	for _, dir := range *dirs {
		files, _ := i.tE.GetFilesInDir(dir)

		for _, file := range *files {
			fileContent, _ := i.fM.GetFileContent(file)
			imports, _ := i.astM.GetImports(*fileContent)

			for _, importPath := range *imports {
				packageList = append(packageList, *domain.NewPackFromPath(importPath))
			}
		}
	}
	return &packageList
}
