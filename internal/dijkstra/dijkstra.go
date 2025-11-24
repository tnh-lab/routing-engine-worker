package dijkstra

import (
	"container/heap"
	"math"
	"github.com/tnh-lab/routing-engine-worker/internal/graph"
)

type Distance map[graph.NodeID]float64

type Previous map[graph.NodeID]graph.NodeID

func Dijsktra(g *graph.Graph, source graph.NodeID) (Distance, Previous) {
	dist := make(Distance)
	prev := make(Previous)

	for node := range g.Adj {
		dist[node] = math.Inf(1)
		prev[node] = -1
	}

	dist[source] = 0

	pq := make(graph.PriorityQueue, 0) 
	heap.Init(&pq)

	heap.Push(&pq, &graph.Item{
		Node: source,
		Dist: 0,
	})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*graph.Item)
		u := item.Node
		
		for _, edge := range g.Adj[u] {
			v := edge.To
			weight := edge.Weight
			alt := dist[u] + weight

			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u

				heap.Push(&pq, &graph.Item{
					Node: v,
					Dist: alt,
				})
			}
		}
	}

	return dist, prev
}
