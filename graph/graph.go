package graph

import "fmt"

type Graph struct {
	vertices []*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make([]*Vertex, 0),
	}

}

func (g Graph) getVertex(key int) *Vertex {

	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}

	return nil
}

func (g *Graph) contains(key int) bool {
	for _, value := range g.vertices {
		if value.key == key {
			return true
		}
	}
	return false
}

func (g *Graph) addVertex(key int) {

	if !g.contains(key) {
		g.vertices = append(g.vertices, NewVertex(key))

	}

}

func (g Graph) String() string {
	var res string

	for _, value := range g.vertices {
		res += fmt.Sprintf("At vertex %d:", value.key)

		for _, vertex := range value.adjacent {
			res += fmt.Sprintf(" %d ", vertex.key)
		}

		res += "\n"
	}

	return res

}

func (g *Graph) addEdge(from, to int) bool {
	fromV := g.getVertex(from)
	toV := g.getVertex(to)

	if fromV == nil || toV == nil || fromV.contains(to) {
		return false
	}

	fromV.adjacent = append(fromV.adjacent, toV)

	return true
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		key: key,
	}
}

func (v *Vertex) contains(key int) bool {
	for _, value := range v.adjacent {
		if value.key == key {
			return true
		}
	}
	return false
}
