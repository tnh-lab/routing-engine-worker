package dijkstra

import (
	"container/heap"
	"math"
	"github.com/tnh-lab/routing-engine-worker/graph"
)

func Dijsktra(g *Graph, source NodeID) (Distance, Previous) {
	dist := make(Distance)
	prev := make(Previous)

	pq := make(PriorityQueue, 0)

}
