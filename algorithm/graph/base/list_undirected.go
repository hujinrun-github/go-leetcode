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
