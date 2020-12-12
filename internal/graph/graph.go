package graph

type Graph struct {
	Paths    [][]int
	vertices int
	edges    map[int][]int
}

func NewGraph(vertices int) *Graph {
	//edges := make([][]int, 0)
	//for i := 0; i < vertices; i++ {
	//	edges[i] = make([]int, 0)
	//}
	return &Graph{
		Paths:    make([][]int, 0),
		vertices: vertices,
		edges:    make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	if _, ok := g.edges[u]; !ok {
		g.edges[u] = make([]int, 0)
	}

	g.edges[u] = append(g.edges[u], v)
}

func (g *Graph) CreatePaths(start, end int) {
	visited := make(map[int]bool)
	pathList := []int{start}

	g.calculatePaths(start, end, visited, pathList)
}

func (g *Graph) calculatePaths(node, end int, visited map[int]bool, path []int) {
	if node == end {
		g.Paths = append(g.Paths, path)
		return
	}

	visited[node] = true

	for _, v := range g.edges[node] {
		if !visited[v] {
			path = append(path, v)
			g.calculatePaths(v, end, visited, path)

			path = remove(path, v)
		}
	}

	visited[node] = false
}

func remove(path []int, node int) []int {
	newPath := make([]int, 0)

	for _, p := range path {
		if p != node {
			newPath = append(newPath, p)
		}
	}

	return newPath
}
