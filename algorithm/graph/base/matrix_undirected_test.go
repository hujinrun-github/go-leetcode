package base

import (
	"fmt"
	"testing"
)

func TestMatrixUndirectedGraph(t *testing.T) {
	//var graph IGraph
	graph := &MatrixUndirectedGraph{}
	graph.InitGraph(5)

	graph.InsertEdge(0, 1, 1)
	graph.InsertEdge(0, 2, 2)
	graph.InsertEdge(1, 3, 1)
	graph.InsertEdge(2, 4, 1)
	graph.InsertEdge(3, 4, 1)
	graph.InsertEdge(3, 4, 2)
	graph.TravelGraphUsingBFS()
}

func TestMatrixUndirectedDijkstraGraph(t *testing.T) {
	//var graph IGraph
	graph := &MatrixUndirectedGraph{}
	graph.InitGraph(6)

	graph.InsertEdge(0, 1, 1)
	graph.InsertEdge(0, 2, 12)
	graph.InsertEdge(1, 2, 9)
	graph.InsertEdge(1, 3, 3)
	graph.InsertEdge(2, 3, 4)
	graph.InsertEdge(2, 4, 5)
	graph.InsertEdge(3, 4, 13)
	graph.InsertEdge(3, 5, 15)
	graph.InsertEdge(4, 5, 4)
	//graph.TravelGraph()
	fmt.Printf("shortest path:%d\n", graph.FindShortestPathUsingDijkstra(0, 5))
}
