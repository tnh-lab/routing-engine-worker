package worker

import (
	"fmt"

	"github.com/tnh-lab/routing-engine-worker/internal/graph"
)

func main() {
	g := graph.NewGraph();
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 2)
	g.AddEdge(0, 2, 4)

	dist, prev := graph.Dijsktra(g, 0)

	fmt.Println("Distances: ", dist)
	fmt.Println("Previous: ", prev)
}