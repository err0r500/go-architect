package json

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/err0r500/go-architect/domain"
)

type D3formatter struct{}

type D3Graph struct {
	Vertices []D3Vertice `json:"nodes"`
	Edges    []D3Edge    `json:"links"`
}

type D3Vertice struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	ID    int    `json:"id"`
}
type D3Edge struct {
	Source int    `json:"source"`
	Target int    `json:"target"`
	Type   string `json: "type"`
}

func (formatter D3formatter) ToJSON(dG *domain.Graph) (string, error) {
	graph := newFromDomain(dG)
	jsonStr, err := json.Marshal(graph)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}

func newFromDomain(dG *domain.Graph) *D3Graph {
	d3G := &D3Graph{}
	rand.Seed(time.Now().Unix())

	fakeCouplings := []string{"low", "high"}

	for _, vertice := range dG.Vertices {
		d3G.Vertices = append(d3G.Vertices, D3Vertice{
			Name:  domain.TrimCurrPackagePathFrom(vertice.Name),
			ID:    vertice.ID,
			Label: vertice.Label,
		})
	}
	for _, edge := range dG.Edges {
		d3G.Edges = append(d3G.Edges, D3Edge{
			Source: edge.Source,
			Target: edge.Target,
			Type:   fakeCouplings[rand.Intn(len(fakeCouplings))],
		})
	}
	return d3G
}
