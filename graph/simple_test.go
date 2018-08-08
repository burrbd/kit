package graph_test

import (
	"testing"

	"github.com/cheekybits/is"

	"github.com/burrbd/kit/graph"
)

func TestNewSimple(t *testing.T) {
	is := is.New(t)
	g := graph.NewSimple()
	is.OK(g)
}

func TestSimple_AddVertex(t *testing.T) {
	is := is.New(t)
	g := graph.NewSimple()
	err := g.AddVertex(vertex{"an_id"})
	is.NoErr(err)
}

func TestSimple_AddVertexDuplicateError(t *testing.T) {
	is := is.New(t)
	g := graph.NewSimple()
	_ = g.AddVertex(vertex{"an_id"})
	err := g.AddVertex(vertex{"an_id"})
	is.Err(err)
}

func TestSimple_AddEdge(t *testing.T) {
	is := is.New(t)
	g := graph.NewSimple()
	v1, v2 := vertex{"id_1"}, vertex{"id_2"}
	_ = g.AddVertex(v1)
	_ = g.AddVertex(v2)
	err := g.AddEdge(v1, v2)
	is.NoErr(err)
}

func TestSimple_AddEdgeMissingVertexError(t *testing.T) {
	is := is.New(t)
	g := graph.NewSimple()
	v1, v2 := vertex{"id_1"}, vertex{"id_2"}
	_ = g.AddVertex(v1)
	err := g.AddEdge(v1, v2)
	is.Err(err)
}

func TestSimple_IsNeighbor(t *testing.T) {
	is := is.New(t)
	g := graph.NewSimple()
	v1, v2 := vertex{"id_1"}, vertex{"id_2"}
	_ = g.AddVertex(v1)
	_ = g.AddVertex(v2)
	_ = g.AddEdge(v1, v2)
	b, err := g.IsNeighbor(v2, v1)
	is.NoErr(err)
	is.True(b)
}

func TestSimple_IsNeighborReturnsFalse(t *testing.T) {
	is := is.New(t)
	g := graph.NewSimple()
	v1, v2 := vertex{"id_1"}, vertex{"id_2"}
	_ = g.AddVertex(v1)
	_ = g.AddVertex(v2)
	b, err := g.IsNeighbor(v2, v1)
	is.NoErr(err)
	is.False(b)
}

func TestSimple_Neighbors(t *testing.T) {
	is := is.New(t)
	g := graph.NewSimple()
	v1, v2 := vertex{"id_1"}, vertex{"id_2"}
	_ = g.AddVertex(v1)
	_ = g.AddVertex(v2)
	_ = g.AddEdge(v1, v2)
	n, err := g.Neighbors(v2)
	is.NoErr(err)
	is.Equal(1, len(n))
}

func TestSimple_NeighborsReturnsNeighbor(t *testing.T) {
	is := is.New(t)
	g := graph.NewSimple()
	v1, v2 := vertex{"id_1"}, vertex{"id_2"}
	_ = g.AddVertex(v1)
	_ = g.AddVertex(v2)
	_ = g.AddEdge(v1, v2)
	n, _ := g.Neighbors(v2)
	is.Equal("id_1", n[0].ID())
}

func TestSimple_NeighborsReturnsManyNeighbors(t *testing.T) {
	is := is.New(t)
	g := graph.NewSimple()
	v1, v2, v3 := vertex{"id_1"}, vertex{"id_2"}, vertex{"id_3"}
	_ = g.AddVertex(v1)
	_ = g.AddVertex(v2)
	_ = g.AddVertex(v3)
	_ = g.AddEdge(v1, v2)
	_ = g.AddEdge(v3, v2)
	n, err := g.Neighbors(v2)
	is.NoErr(err)
	is.Equal(2, len(n))

}

func TestSimple_NeighborsReturnsManyCorrectNeighbors(t *testing.T) {
	is := is.New(t)
	g := graph.NewSimple()
	v1, v2, v3, v4 := vertex{"id_1"}, vertex{"id_2"}, vertex{"id_3"}, vertex{"id_4"}
	_ = g.AddVertex(v1)
	_ = g.AddVertex(v2)
	_ = g.AddVertex(v3)
	_ = g.AddVertex(v4)
	_ = g.AddEdge(v1, v2)
	_ = g.AddEdge(v3, v2)
	vertices, _ := g.Neighbors(v2)
	is.Equal(2, len(vertices))
	is.True(contains(v1, vertices))
	is.True(contains(v3, vertices))
	is.False(contains(v4, vertices))
}

type vertex struct {
	id string
}

func (v vertex) ID() string {
	return v.id
}

func contains(needle vertex, haystack []graph.Identifier) bool {
	for _, v := range haystack {
		if needle.ID() == v.ID() {
			return true
		}
	}
	return false
}
