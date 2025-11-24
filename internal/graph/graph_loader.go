package graph

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
)

func LoadGraphFromCSV(path string) (*Graph, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open CSV: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("cannot read header: %w", err)
	}

	g := NewGraph()
	
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		from, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("invalid 'from': %w", err)
		}
		
		to, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, fmt.Errorf("invalid 'to': %w", err)
		}

		weight, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid 'weight': %w", err)
		}

		g.AddEdge(NodeID(from), NodeID(to), weight)
		g.AddEdge(NodeID(to), NodeID(from), weight)		
	}

	return g, nil
}