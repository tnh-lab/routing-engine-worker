package graph

import (
	"encoding/binary"
	"fmt"
	"os"
	"unsafe"
)

func LoadCSR(path string) (*CSRGraph, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    if len(data) < 8 {
        return nil, fmt.Errorf("file too small: %s", path)
    }

    // Header
    N := binary.LittleEndian.Uint32(data[0:4])
    M := binary.LittleEndian.Uint32(data[4:8])

    offsetsSize := (N + 1) * 4
    edgesSize   := M * 4
    weightsSize := M * 4

    needed := 8 + offsetsSize + edgesSize + weightsSize
    if uint32(len(data)) < needed {
        return nil, fmt.Errorf("CSR file truncated: expected %d bytes, got %d", needed, len(data))
    }

    offStart := uint32(8)
    edgeStart := offStart + offsetsSize
    weightStart := edgeStart + edgesSize

    g := &CSRGraph{
        NumNodes: N,
        NumEdges: M,
        Offsets:  make([]uint32, N+1),
        Edges:    make([]uint32, M),
        Weights:  make([]float32, M),
    }

    // Load offsets[]
    for i := uint32(0); i < N+1; i++ {
        g.Offsets[i] = binary.LittleEndian.Uint32(data[offStart+(i*4):])
    }

    // Load edges[]
    for i := uint32(0); i < M; i++ {
        g.Edges[i] = binary.LittleEndian.Uint32(data[edgeStart+(i*4):])
    }

    // Load weights[]
    // reinterpret bytes as float32
    for i := uint32(0); i < M; i++ {
        bits := binary.LittleEndian.Uint32(data[weightStart+(i*4):])
        g.Weights[i] = float32FromBits(bits)
    }

    return g, nil
}

func float32FromBits(b uint32) float32 {
    return *(*float32)(unsafe.Pointer(&b))
}

func (g *CSRGraph) Neighbors(u uint32) (start, end uint32) {
    return g.Offsets[u], g.Offsets[u+1]
}

