package dijkstra

import (
    "container/heap"
    "math"

    "github.com/tnh-lab/routing-engine-worker/internal/graph"
)

func CSR(g *graph.CSRGraph, source uint32) ([]float32, []uint32) {
    N := g.NumNodes

    // Distances and predecessor arrays
    dist := make([]float32, N)
    prev := make([]uint32, N)

    for i := uint32(0); i < N; i++ {
        dist[i] = float32(math.Inf(1))
        prev[i] = ^uint32(0) // 0xFFFFFFFF = no parent
    }
    dist[source] = 0

    pq := &graph.PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, &graph.Item{Node: source, Dist: 0})

    for pq.Len() > 0 {
        item := heap.Pop(pq).(*graph.Item)
        u := item.Node

        start, end := g.Neighbors(u)
        for i := start; i < end; i++ {
            v := g.Edges[i]
            w := g.Weights[i]

            alt := dist[u] + w
            if alt < dist[v] {
                dist[v] = alt
                prev[v] = u
                heap.Push(pq, &graph.Item{Node: v, Dist: alt})
            }
        }
    }

    return dist, prev
}
