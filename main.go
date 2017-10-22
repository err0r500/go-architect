package main

import (
	"log"

	"github.com/err0r500/go-architect/domain"
	AstM "github.com/err0r500/go-architect/interfaces/astManager"
	FM "github.com/err0r500/go-architect/interfaces/fileManager"
	TE "github.com/err0r500/go-architect/interfaces/treeExplorer"
)

type ImportsFinderInteractor struct {
	tE   TreeExplorer
	fM   FileManager
	astM AstManager
}

func main() {

	dummy := ImportsFinderInteractor{
		tE:   TE.TreeExplorer{},
		fM:   FM.FileManager{},
		astM: AstM.AstManager{},
	}

	dummy.GetAllImports()

}

// juste un gros bloc pour montrer l'idée initiale, surement naîve
func (i ImportsFinderInteractor) GetAllImports() *[]domain.Pack {
	dirs, _ := i.tE.GetDirsInTree(".")
	packageList := []domain.Pack{}

	for _, dir := range *dirs {
		fPathes, _ := i.tE.GetFilesInDir(dir)

		for _, fPath := range *fPathes {
			f := domain.File{Path: dir + "/" + fPath}
			fileContent, err := i.fM.GetFileContent(f)
			if err != nil {
				log.Print(err)
				continue
			}
			imports, _ := i.astM.GetImports(*fileContent)

			for _, importPath := range *imports {
				packageList = append(packageList, *domain.NewPackFromPath(importPath))
			}
			logThisDep(f, *imports)
		}
	}
	return &packageList
}

// just for tests
func logThisDep(f domain.File, imports []string) {
	logStr := ""
	if len(imports) == 0 {
		log.Print(f.GetPath(), "depends on nothing")
		return
	}

	logStr = f.GetPath() + " depends on :\n"
	for _, imp := range imports {
		p := domain.NewPackFromPath(imp)
		logStr += "	  " + p.String() + "\n"
	}
	log.Print(logStr)
}
