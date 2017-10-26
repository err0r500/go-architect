# go-architect
[![BCH compliance](https://bettercodehub.com/edge/badge/Err0r500/go-architect?branch=master)](https://bettercodehub.com/)
[![Build Status](http://35.202.103.169:8000/api/badges/Err0r500/go-architect/status.svg)](http://35.202.103.169:8000/Err0r500/go-architect)

![ScreenShot](https://raw.githubusercontent.com/Err0r500/go-architect/master/doc/screenshots/wide.png)

## Usage

### Install 
```
go get github.com/err0r500/go-architect
cd $GOPATH/src/github.com/err0r500/go-architect/
go install
```

### Usage
```
cd $GOPATH/src/whatever/folder/you/want/to/analyze
$GOPATH/bin/go-architect
```
Navigate to  [http://localhost/8080](http://localhost/8080)


## Roadmaps
### v0.1
- [x] find all imports in a directory tree
- [x] filter out the duplicates
- [x] build the corresponding [directed graph](https://en.wikipedia.org/wiki/Directed_graph) 
- [x] format the graph for a visualization lib or 2 (see below)
- [x] visualize the graph (see below)

##{} v0.2
- [ ] add simple qualification / weight to vertices from default settings
- [ ] Desktop version

### v0.3
- [ ] dynamically add & customize rules from UI (grapQL ?) 
- [ ] add new rules as plugins ?

## Tech ideas :
- Vizualisation : 
    - [d3js](https://github.com/d3/d3/wiki/Gallery)
    - [cytoscape.js](https://github.com/cytoscape/cytoscape.js)

- Desktop :
    - [electron](https://electron.atom.io/)
    
- Storage : 
    - [ArangoDB](https://www.arangodb.com/)

En regardant ce que des libs JS ci-dessus font à partir d'un simple JSON, l'étape DB peut sans doute se faire dans un second temps... On peut déjà faire des essais avec l'interface JSONwriter qui génère un fichier formatté pour une lib ou une autre (en fct de quel struct l'implémente) et on plug ça dans un exemple de démo ...

