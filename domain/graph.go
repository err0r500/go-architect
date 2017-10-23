package domain

type Graph struct {
	Vertices []Vertice
	Edges    []Edge
}

type Vertice struct {
	Pack
	Name  string
	Label string
	ID    int
}

type Edge struct {
	Source int
	Target int
}

func (graph *Graph) AppendRootNode(p *Pack) int {
	for i, vertice := range graph.Vertices {
		if vertice.Pack.packagePath == p.packagePath {
			return i
		}
	}

	insertID := len(graph.Vertices)
	vertice := Vertice{Pack: *p, Name: p.GetPath(), Label: p.GetClass(), ID: insertID}
	graph.Vertices = append(graph.Vertices, vertice)
	return insertID
}

func (graph *Graph) BuildGraph(verticeID int, p *Pack) {
	nodeID := graph.AppendRootNode(p)
	if verticeID != nodeID {
		graph.Edges = append(graph.Edges, Edge{Source: verticeID, Target: nodeID})
	}
}
