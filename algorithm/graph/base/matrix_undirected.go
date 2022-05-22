package base

import "fmt"

// this file defined a struct that represent graph using adjacency matrix

type MatrixUndirectedGraph struct {
	nv int     // the number of vertices
	ne int     // the number of edges
	g  [][]int // the representation of graph using adjacency matrix
}

func (m *MatrixUndirectedGraph) InitGraph(nv int) {
	m.nv = nv
	m.g = make([][]int, nv)
	for i := 0; i < nv; i++ {
		m.g[i] = make([]int, nv)
		for j := 0; j < nv; j++ {
			m.g[i][j] = -1 // -1 represent the distance is infinity
		}
	}
}

func (m *MatrixUndirectedGraph) InsertEdge(u, v, weight int) error {
	if u >= m.nv || u < 0 || v >= m.nv || v < 0 {
		return fmt.Errorf("vertice is not illegal, v:%v, u:%v", u, v)
	}

	if m.g[u][v] != -1 || m.g[v][u] != -1 {
		return fmt.Errorf("this edge has been inserted")
	}

	m.g[u][v] = weight
	m.g[v][u] = weight
	m.ne++
	return nil
}

func (m *MatrixUndirectedGraph) TravelGraph() {
	// the travel just show the adjacency vertices of each vertice
	for i := 0; i < m.nv; i++ {
		fmt.Printf("vertice:%d -> adjacency vertices:[", i)
		for j := 0; j < m.nv; j++ {
			if j != i {
				fmt.Printf("vertice:%d,weight:%d;", j, m.g[i][j])
			}
		}
		fmt.Println()
	}
}

func (m *MatrixUndirectedGraph) GetVerticeNumber() int {
	return m.nv
}

func (m *MatrixUndirectedGraph) GetEdgeNumber() int {
	return m.ne
}
