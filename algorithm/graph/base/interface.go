package base

// the file defined a common interface that giving all operations of graph

type IGraph interface {
	InitGraph(int)
	InsertEdge(int, int, int) error
	GetVerticeNumber() int
	GetEdgeNumber() int
	TravelGraph()
	TravelGraphUsingDFS()
	TravelGraphUsingBFS()
	FindShortestPathUsingDijkstra(int, int) int
}
