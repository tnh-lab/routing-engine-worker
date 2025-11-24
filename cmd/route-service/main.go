package main

import (
	"context"
	"log"
	"math"
	"net"
	"os"

	"google.golang.org/grpc"
	"github.com/tnh-lab/routing-engine-worker/internal/dijkstra"
	"github.com/tnh-lab/routing-engine-worker/internal/graph"
	pb "github.com/tnh-lab/routing-engine-worker/pb"
)

type server struct {
	pb.UnimplementedRoutingServiceServer
	g *graph.CSRGraph
}

func NewServer(g *graph.CSRGraph) *server {
	return &server{g: g}
}

func (s *server) ShortestPath(ctx context.Context, req *pb.ShortestPathRequest) (*pb.ShortestPathResponse, error) {
	source := uint32(req.GetSource())
	target := uint32(req.GetTarget())

	// Run Dijkstra on CSR
	dist, prev := dijkstra.CSR(s.g, source)

	// unreachable?
	if dist[target] == float32(math.Inf(1)) {
		return &pb.ShortestPathResponse{
			Path:          nil,
			TotalDistance: 0,
		}, nil
	}

	// reconstruct path (target → source, reversed)
	pathNodes := []uint32{}
	cur := target
	noPrev := uint32(0xFFFFFFFF)

	for cur != noPrev {
		pathNodes = append([]uint32{cur}, pathNodes...)
		if cur == source {
			break
		}
		cur = prev[cur]
	}

	resp := &pb.ShortestPathResponse{
		TotalDistance: float64(dist[target]),
	}

	// convert to protobuf format
	for _, n := range pathNodes {
		resp.Path = append(resp.Path, &pb.PathNode{
			NodeId: int64(n),
			Dist:   float64(dist[n]),
		})
	}

	return resp, nil
}

func main() {
	csrPath := os.Getenv("GRAPH_CSR")
	if csrPath == "" {
		csrPath = "data/roadNet-CA.csr"
	}

	// Load CSR graph
	g, err := graph.LoadCSR(csrPath)
	if err != nil {
		log.Fatalf("failed to load CSR graph: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRoutingServiceServer(grpcServer, NewServer(g))

	log.Println("RoutingService gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
