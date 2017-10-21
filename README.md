# go-architect
[![BCH compliance](https://bettercodehub.com/edge/badge/Err0r500/go-architect?branch=master)](https://bettercodehub.com/)
[![Build Status](http://35.202.103.169:8000/api/badges/Err0r500/go-architect/status.svg)](http://35.202.103.169:8000/Err0r500/go-architect)


## Roadmap v0.1
- [ ] find all imports in a directory tree
- [ ] filter out the duplicates
- [ ] build the corresponding [directed graph](https://en.wikipedia.org/wiki/Directed_graph) 
- [ ] format the graph for a visualization lib or 2 (see below)
- [ ] visualize the graph (see below)
- [ ] build for Linux & Windows

## Roadmap v0.2
- [ ] add simple qualification / weight to vertices from default settings

## Roadmap v0.3
- [ ] dynamically add & customize rules from UI (grapQL ?) 
- [ ] add new rules as plugins ?

## Tech ideas :
- Vizualisation : 
    - [d3js](https://github.com/d3/d3/wiki/Gallery)
    - [cytoscape.js](https://github.com/cytoscape/cytoscape.js)

- Storage : 
    - [ArangoDB](https://www.arangodb.com/)

En regardant ce que des libs JS ci-dessus font à partir d'un simple JSON, l'étape DB peut sans doute se faire dans un second temps... On peut déjà faire des essais avec l'interface JSONwriter qui génère un fichier formatté pour une lib ou une autre (en fct de quel struct l'implémente) et on plug ça dans un exemple de démo ...

