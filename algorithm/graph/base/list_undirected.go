package base

import "fmt"

// this file defined a struct that represent graph using adjacency linked list

type vnode struct {
	node int    // the vertice value
	next *vnode // the adjacency vertice of this vertice
}

type ListUndirectedGraph struct {
	nv        int                 // the number of vertices
	ne        int                 // the number of edges
	weightMap map[int]map[int]int // record the weight of edges
	g         []vnode             // the representation of graph using adjacency linked list
}

func (l *ListUndirectedGraph) InitGraph(nv int) {
	l.nv = nv
	l.g = make([]vnode, nv)
	for i := 0; i < nv; i++ {
		l.g[i].node = i
	}
}

func (l *ListUndirectedGraph) InsertEdge(u, v, weight int) error {
	if u >= l.nv || u < 0 || v >= l.nv || v < 0 {
		return fmt.Errorf("vertice is not illegal, v:%v, u:%v", u, v)
	}

	uLastNode := &l.g[u]
	for uLastNode.next != nil {
		if uLastNode.next.node == v {
			return fmt.Errorf("the edge between vertice u and v is already exist")
		}
		uLastNode = uLastNode.next
	}

	vLastNode := &l.g[v]
	for vLastNode.next != nil {
		if vLastNode.next.node == u {
			return fmt.Errorf("the edge between vertice u and v is already exist")
		}
		vLastNode = vLastNode.next
	}

	uLastNode.next = &vnode{node: v}
	vLastNode.next = &vnode{node: u}
	if l.weightMap == nil {
		l.weightMap = map[int]map[int]int{}
	}

	if l.weightMap[v] == nil {
		l.weightMap[v] = map[int]int{}
	}

	if l.weightMap[u] == nil {
		l.weightMap[u] = map[int]int{}
	}

	l.weightMap[v][u] = weight
	l.weightMap[u][v] = weight

	l.ne++
	return nil
}

func (l *ListUndirectedGraph) TravelGraph() {
	// the travel just show the adjacency vertices of each vertice
	for i := 0; i < l.nv; i++ {
		fmt.Printf("vertice:%d -> adjacency vertices:[", i)
		next := l.g[i].next
		for next != nil {
			fmt.Printf("vertice:%d,weight:%d;", next.node, l.weightMap[i][next.node])
			next = next.next
		}

		fmt.Println("]")
	}
}

func (l *ListUndirectedGraph) GetVerticeNumber() int {
	return l.nv
}

func (l *ListUndirectedGraph) GetEdgeNumber() int {
	return l.ne
}

func (l *ListUndirectedGraph) TravelGraphUsingDFS() {
	result := []int{}
	visited := map[int]struct{}{}
	var dfs func(int)
	dfs = func(i int) {
		visited[i] = struct{}{}
		result = append(result, i)
		next := l.g[i].next
		for next != nil {
			node := next.node
			if _, contain := visited[node]; !contain {
				dfs(node)

			}
			next = next.next
		}
	}

	for i := 0; i < l.nv; i++ {
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

func (l *ListUndirectedGraph) TravelGraphUsingBFS() {
	result := []int{}
	visited := map[int]struct{}{}
	bfs := func(i int) {
		// visited[i] = struct{}{}
		// result = append(result, i)
		queue := []int{i}
		qlen := 1
		for qlen > 0 {
			tmp := queue[0]
			queue = queue[1:]
			visited[tmp] = struct{}{}
			result = append(result, tmp)
			// travel
			next := l.g[tmp].next
			for next != nil {
				node := next.node
				if _, contain := visited[node]; !contain {
					visited[node] = struct{}{}
					queue = append(queue, node)
				}
				next = next.next
			}

			qlen = len(queue)
		}
	}

	for i := 0; i < l.nv; i++ {
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

func (l *ListUndirectedGraph) FindShortestPathUsingDijkstra(start, end int) int {
	book := map[int]struct{}{}
	dist := make([]int, l.nv)
	// initialize the shortest path
	for i := 0; i < l.nv; i++ {
		if i != start {
			dist[i] = INT_MAX
		}
	}
	adj := l.g[start].next
	for adj != nil {
		dist[adj.node] = l.weightMap[start][adj.node]
		adj = adj.next
	}

	book[start] = struct{}{}
	for i := 1; i < l.nv; i++ {
		shortestNode := -1
		shortestPath := INT_MAX
		// find the shortest path
		for i := 0; i < l.nv; i++ {
			if _, contain := book[i]; !contain && dist[i] < shortestPath {
				shortestPath = dist[i]
				shortestNode = i
			}
		}

		book[shortestNode] = struct{}{}
		// update 'dist'
		adj := l.g[shortestNode].next
		for adj != nil {
			if _, contain := book[adj.node]; !contain && dist[shortestNode]+l.weightMap[shortestNode][adj.node] < dist[adj.node] {
				dist[adj.node] = dist[shortestNode] + l.weightMap[shortestNode][adj.node]
			}
			adj = adj.next
		}
	}
	return dist[end]
}

func (l *ListUndirectedGraph) IsCircleUsingTopologySort() bool {
	// step1: calculate the degree of each node
	// step2: add the node which has the number of degree less than or equal one to queue
	degressList := make([]int, l.nv)
	queue := []int{}

	for i := 0; i < l.nv; i++ {
		count := 0
		adj := l.g[i].next
		for adj != nil {
			count++
			adj = adj.next
		}

		degressList[i] = count

		if count <= 1 {
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
		adj := l.g[tmpNode].next
		for adj != nil {
			degressList[adj.node]--
			if _, ok := traveledMap[adj.node]; !ok && degressList[adj.node] <= 1 {
				queue = append(queue, adj.node)
			}
			adj = adj.next
		}
	}

	// step4: check the number of all traveled node whether it is m.nv
	return len(traveledMap) != l.nv
}
