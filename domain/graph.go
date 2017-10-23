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
	vertice := Vertice{Pack: *p, Name: p.GetPath(), Label: p.GetClass(), ID: len(graph.Vertices)}

	graph.Vertices = append(graph.Vertices, vertice)
	return len(graph.Vertices)
}

func (graph *Graph) BuildGraph(verticeID int, p *Pack) {
	nodeID := graph.AppendRootNode(p)
	graph.Edges = append(graph.Edges, Edge{Source: verticeID, Target: nodeID})
	return
}
