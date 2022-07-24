package graph

import (
	"testing"
)

func initGraph() *Graph {
	g := NewGraph()

	for i := 0; i < 5; i++ {
		g.addVertex(i)
	}

	return g
}

func TestContains(t *testing.T) {
	g := initGraph()

	if !g.contains(1) {
		t.Errorf("Graph contains was incorrect, got: %t, want: %t.", g.contains(1), true)
	}

}

func TestAddVertex(t *testing.T) {
	g := initGraph()

	g.addVertex(21)

	if !g.contains(21) {
		t.Errorf("Graph add vertex was incorrect, got: %t, want: %t.", g.contains(21), true)
	}

}

func TestAddVertex2(t *testing.T) {
	g := initGraph()

	size := len(g.vertices)

	g.addVertex(1)

	if len(g.vertices) != size {
		t.Errorf("Graph size vertices was incorrect, got: %d, want: %d.", len(g.vertices), size)

	}

}

func TestAddEddEdge(t *testing.T) {
	g := initGraph()

	g.addEdge(1, 2)

	if len(g.vertices[1].adjacent) == 0 {
		t.Errorf("Graph add edge was incorrect, got: %d, want: %d.", len(g.vertices[1].adjacent), 0)

	}

}

func TestAddEddEdge2(t *testing.T) {
	g := initGraph()

	g.addEdge(10, 200)

	if g.addEdge(10, 200) {
		t.Errorf("Graph add edge was incorrect, got: %t, want: %t.", g.addEdge(10, 200), false)

	}

}

