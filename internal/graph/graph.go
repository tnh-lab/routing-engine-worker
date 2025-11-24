package graph

type NodeID int

type Edge struct {
	To NodeID
	Weight float64
}

type Graph struct {
	Adj map[NodeID][]Edge
}

func NewGraph() *Graph {
	g := &Graph{
		Adj: make(map[NodeID][]Edge),
	}

	return g
}

func (g *Graph) AddEdge(from NodeID, to NodeID, weight float64) {
	g.Adj[from] = append(g.Adj[from], Edge{To: to, Weight: weight})

	if _, exists := g.Adj[to]; !exists {
		g.Adj[to] = []Edge{}
	}
}