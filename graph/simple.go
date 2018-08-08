// Package graph provides a simple graph data structure. Motivation for this
// package was to support a Go implementation of a well-known board game.
package graph

import (
	"fmt"
	"sync"
)

// Identifier has an ID method.
type Identifier interface {
	ID() string
}

// Simple is an implementation of an undirected unweighted graph.
type Simple struct {
	sync.RWMutex
	vertices []Identifier
	edges    map[string][]Identifier
}

// NewSimple initialises a new Simple struct.
func NewSimple() *Simple {
	return &Simple{
		vertices: make([]Identifier, 0),
		edges:    make(map[string][]Identifier),
	}
}

// AddVertex adds a new vertex to the Simple and returns an error
// when a vertex with the same ID already exists.
func (g *Simple) AddVertex(v Identifier) error {
	g.Lock()
	defer g.Unlock()
	if g.hasVertex(v) {
		return fmt.Errorf("vertex id '%s' exists", v.ID())
	}
	g.vertices = append(g.vertices, v)
	return nil
}

// AddEdge adds an edge to the graph and returns an error if the
// vertices have not already been added.
func (g *Simple) AddEdge(v1, v2 Identifier) error {
	g.Lock()
	defer g.Unlock()
	if !g.hasVertex(v1) || !g.hasVertex(v2) {
		return fmt.Errorf("vertex id '%s' or '%s' missing", v1.ID(), v2.ID())
	}
	if _, ok := g.edges[v1.ID()]; !ok {
		g.edges[v1.ID()] = make([]Identifier, 0)
	}
	if _, ok := g.edges[v2.ID()]; !ok {
		g.edges[v2.ID()] = make([]Identifier, 0)
	}
	g.edges[v1.ID()] = append(g.edges[v1.ID()], v2)
	g.edges[v2.ID()] = append(g.edges[v2.ID()], v1)
	return nil
}

// IsNeighbor determines if vertices v1 and v2 form an edge. Returns
// an error if one or more of the vertices do not exist in the graph.
func (g *Simple) IsNeighbor(v1, v2 Identifier) (bool, error) {
	g.RLock()
	defer g.RUnlock()
	if !g.hasVertex(v1) || !g.hasVertex(v2) {
		return false, fmt.Errorf("vertex id '%s' or '%s' missing", v1.ID(), v2.ID())
	}
	return g.hasEdge(v1, v2), nil
}

// Neighbors returns all vertices that share an edge with Vertex v.
func (g *Simple) Neighbors(v Identifier) ([]Identifier, error) {
	g.RLock()
	defer g.RUnlock()
	if !g.hasVertex(v) {
		return nil, fmt.Errorf("vertex id '%s' missing", v.ID())
	}
	if g.edges[v.ID()] == nil {
		g.edges[v.ID()] = make([]Identifier, 0)
	}
	return g.edges[v.ID()], nil
}

func (g *Simple) hasVertex(v Identifier) bool {
	for _, vv := range g.vertices {
		if vv.ID() == v.ID() {
			return true
		}
	}
	return false
}

func (g *Simple) hasEdge(v1, v2 Identifier) bool {
	for _, v := range g.edges[v1.ID()] {
		if v.ID() == v2.ID() {
			return true
		}
	}
	return false
}
