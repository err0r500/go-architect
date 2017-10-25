package main

import (
	"go/build"
	"io"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/err0r500/go-architect/domain"
	"github.com/err0r500/go-architect/interfaces"
	AstM "github.com/err0r500/go-architect/interfaces/astManager"
	FM "github.com/err0r500/go-architect/interfaces/fileManager"
	"github.com/err0r500/go-architect/interfaces/json"
	TE "github.com/err0r500/go-architect/interfaces/treeExplorer"
)

type ImportsFinderInteractor struct {
	tE    interfaces.TreeExplorer
	fM    interfaces.FileManager
	astM  interfaces.AstManager
	jsonW interfaces.JSONwriter
}

func main() {

	i := ImportsFinderInteractor{
		tE:    TE.TreeExplorer{},
		fM:    FM.FileManager{},
		astM:  AstM.AstManager{},
		jsonW: json.D3formatter{},
	}

	// graph := i.GetAllImports()
	// jsonPayload, _ := i.jsonW.ToJSON(graph)
	// err := i.fM.Write(domain.File{Path: "./js/testGraph.json", Content: []byte(jsonPayload)})

	_, filename, _, _ := runtime.Caller(0)
	sourceDir := path.Dir(filename)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, req, sourceDir+"/js/d3_1/index.html")
	})

	http.HandleFunc("/vizu.js", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, req, sourceDir+"/js/d3_1/vizu.js")

	})

	http.HandleFunc("/data/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		graph := i.GetAllImports("./")
		jsonPayload, _ := i.jsonW.ToJSON(graph)
		io.WriteString(w, jsonPayload)
	})

	http.ListenAndServe(":8080", nil)

}

// juste un gros bloc pour montrer l'idée initiale, surement naîve
func (i ImportsFinderInteractor) GetAllImports(path string) *domain.Graph {

	dirs, _ := i.tE.GetDirsInTree(filepath.Dir(path))
	graph := &domain.Graph{}

	for _, dir := range *dirs {
		dir, err := filepath.Abs(dir)
		if err != nil {
			continue
		}

		currVerticeID := appendVertice(graph, dir)
		fPathes, _ := i.tE.GetFilesInDir(dir)

		for _, fPath := range *fPathes {
			f := domain.File{Path: dir + "/" + fPath}

			imports, err := i.astM.GetImportsFromFile(f.GetPath())
			if err != nil {
				continue
			}

			for _, importPath := range *imports {
				graph.BuildGraph(currVerticeID, domain.NewPackFromPath(importPath))
			}
		}
	}
	return graph
}

func appendVertice(graph *domain.Graph, srcFolder string) int {
	src, _ := filepath.Abs(srcFolder)
	src = strings.Replace(src, build.Default.GOPATH+"/src/", "", -1)
	return graph.AppendRootNode(domain.NewPackFromPath(src))
}
