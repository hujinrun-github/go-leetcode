package base

import (
	"fmt"
)

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
			if i != j {
				m.g[i][j] = -1 // -1 represent the distance is infinity
			}
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

// travel the graph using dfs
func (m *MatrixUndirectedGraph) TravelGraphUsingDFS() {
	result := []int{}
	visited := map[int]struct{}{}
	var dfs func(int)
	dfs = func(i int) {
		visited[i] = struct{}{}
		result = append(result, i)
		for j := 0; j < m.nv; j++ {
			if _, contain := visited[j]; !contain && m.g[i][j] > 0 {
				dfs(j)
			}
		}
	}

	for i := 0; i < m.nv; i++ {
		if _, contain := visited[i]; !contain {
			dfs(i)
			fmt.Print("[")
			for _, v := range result {
				fmt.Printf("\t%d", v)
			}
			fmt.Print("]")
			result = []int{}
		}
	}
}

// travel the graph using bfs
func (m *MatrixUndirectedGraph) TravelGraphUsingBFS() {
	result := []int{}
	visited := map[int]struct{}{}
	bfs := func(i int) {
		// result = append(result, i)
		// visited[i] = struct{}{}
		queue := []int{i}
		qlen := 1
		for qlen > 0 {
			tmp := queue[0]
			queue = queue[1:]
			result = append(result, tmp)
			visited[tmp] = struct{}{}
			for j := 0; j < m.nv; j++ {
				if _, contain := visited[j]; !contain && m.g[tmp][j] > 0 {
					visited[j] = struct{}{}
					queue = append(queue, j)
				}
			}
			qlen = len(queue)
		}
	}

	for i := 0; i < m.nv; i++ {
		if _, contain := visited[i]; !contain {
			bfs(i)
			fmt.Print("[")
			for _, v := range result {
				fmt.Printf("\t%d", v)
			}
			fmt.Print("]")
			result = []int{}
		}
	}
}

func (m *MatrixUndirectedGraph) FindShortestPathUsingDijkstra(start, end int) int {
	book := map[int]struct{}{}
	dist := make([]int, m.nv)
	// initialize the shortest path
	for i := 0; i < m.nv; i++ {
		if i != start {
			if m.g[start][i] == -1 {
				dist[i] = INT_MAX
			} else {
				dist[i] = m.g[start][i]
			}
		}
	}
	//fmt.Println(dist)

	// find the shortest path
	book[start] = struct{}{}
	for len(book) != m.nv {
		shortestNode := -1
		shortestPath := INT_MAX
		// find the shortest path
		for i := 0; i < m.nv; i++ {
			if _, contain := book[i]; !contain && dist[i] < shortestPath {
				shortestPath = dist[i]
				shortestNode = i
			}
		}

		fmt.Println(dist)
		book[shortestNode] = struct{}{}

		// update 'dist'
		for i := 0; i < m.nv; i++ {
			if _, contain := book[i]; !contain && m.g[shortestNode][i] != -1 && dist[shortestNode]+m.g[shortestNode][i] < dist[i] {
				dist[i] = dist[shortestNode] + m.g[shortestNode][i]
			}
		}
	}
	return dist[end]
}

func (m *MatrixUndirectedGraph) IsCircleUsingTopologySort() bool {
	// step1 calculate the degree of each node
	degressList := make([]int, m.nv)

	for i := 0; i < m.nv; i++ {
		count := 0
		for j := 0; j < m.nv; j++ {
			if i != j && m.g[i][j] != -1 {
				count++
			}
		}

		degressList[i] = count
	}

	// judge whether the graph is circle by topology sort
	// step2: add the node which has the number of degree less than or equal one to queue
	queue := []int{}
	for i := 0; i < m.nv; i++ {
		if degressList[i] <= 1 {
			queue = append(queue, i)
		}
	}

	// step3: travel queue and pop the first element, and reduce the degree of adjacent nodes by one,
	//		  if the degree of adjacent nodes are less than or equal one, add them to queue
	traveledMap := map[int]struct{}{}
	for len(queue) > 0 {
		tmpNode := queue[0]
		queue = queue[1:]
		traveledMap[tmpNode] = struct{}{}
		// find the adjacent
		for i := 0; i < m.nv; i++ {
			if i != tmpNode && m.g[tmpNode][i] != 0 {
				degressList[i]--
			}

			if _, ok := traveledMap[i]; !ok && degressList[i] <= 1 {
				queue = append(queue, i)
			}
		}
	}

	// step4: check the number of all traveled node whether it is m.nv
	return len(traveledMap) != m.nv
}

func (m *MatrixUndirectedGraph) IsCircleUsingDFS() bool {
	return true
}
