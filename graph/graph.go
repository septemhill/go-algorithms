package graph

import (
	"errors"
	"fmt"
)

type Edge struct {
	Start int
	End   int
}

type Graph struct {
	Vertex int
	Edges  []Edge
}

func (g *Graph) AddEdge(start, end int) error {
	if start >= g.Vertex || end >= g.Vertex || start < 0 || end < 0 {
		return errors.New("Start/End pointer over/under " + fmt.Sprintf("%d ~ %d", 0, g.Vertex))
	}
	g.Edges = append(g.Edges, Edge{start, end})

	return nil
}

func (g Graph) EdgeList() {
	for _, e := range g.Edges {
		fmt.Println(e)
	}
}

func NewGraph(vertex int) Graph {
	return Graph{vertex, []Edge{}}
}
