package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/tnh-lab/routing-engine-worker/internal/dijkstra"
	"github.com/tnh-lab/routing-engine-worker/internal/graph"
	"github.com/tnh-lab/routing-engine-worker/pb"
)

type server struct {
	pb.UnimplementedRoutingServiceServer
	g *graph.Graph
}

func (s *server) ShortestPath(ctx context.Context, req *pb.ShortestPathRequest) (*pb.ShortestPathResponse, error) {
	dist, prev := dijkstra.Dijsktra(s.g, graph.NodeID(req.Source))

	path := []int64{}
	current := req.Target

	for current != -1 {
		path = append([]int64{current}, path...)
		current = int64(prev[graph.NodeID(current)])
	}

	resp := &pb.ShortestPathResponse{
		TotalDistance: dist[graph.NodeID(req.Target)],
	}

	for _, node := range path {
		resp.Path = append(resp.Path, &pb.PathNode{
			NodeId: node,
			Dist: dist[graph.NodeID(node)],
		})
	}

	return resp, nil

}

func main() {
    // Load or build your graph
    g, err := graph.LoadGraphFromCSV("data/edges.csv")
	if err != nil {
        log.Fatalf("Failed to load graph: %v", err)
    }
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatal(err)
    }

    s := grpc.NewServer()
    pb.RegisterRoutingServiceServer(s, &server{g: g})

    log.Println("Routing gRPC server listening on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatal(err)
    }
}