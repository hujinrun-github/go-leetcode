package base

import (
	"testing"
)

func TestMatrixUndirectedGraph(t *testing.T) {
	var graph IGraph
	graph = &MatrixUndirectedGraph{}
	graph.InitGraph(5)

	graph.InsertEdge(0, 1, 1)
	graph.InsertEdge(0, 2, 2)
	graph.InsertEdge(1, 3, 1)
	graph.InsertEdge(2, 4, 1)
	graph.InsertEdge(3, 4, 1)
	graph.InsertEdge(3, 4, 2)
	graph.TravelGraphUsingBFS()
}
