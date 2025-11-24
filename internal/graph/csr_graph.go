package graph

type CSRGraph struct {
    NumNodes uint32
    NumEdges uint32

    // Offsets has length NumNodes+1
    Offsets []uint32

    // Edges has length NumEdges
    Edges   []uint32

    // Weights has length NumEdges (float32)
    Weights []float32
}
