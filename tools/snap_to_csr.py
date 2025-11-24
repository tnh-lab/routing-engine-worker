import struct
import numpy as np
import pathlib


INPUT = "tools/roadNet-CA.txt"
OUTPUT = "tools/roadNet-CA.csr"

edges = []
max_node = 0

with open(INPUT) as f:
    for line in f:
        if line.startswith("#"):
            continue
        u, v = map(int, line.split())
        edges.append((u, v, 1.0))
        edges.append((v, u, 1.0))
        max_node = max(max_node, u, v)

N = max_node + 1

edges.sort(key=lambda x: x[0])

offsets = np.zeros(N+1, dtype=np.uint32)
neighbors = np.zeros(len(edges), dtype=np.uint32)
weights = np.zeros(len(edges), dtype=np.float32)

cur = 0
for i, (u, v, w) in enumerate(edges):
    while cur <= u:
        offsets[cur] = i
        cur += 1
    neighbors[i] = v
    weights[i] = w

offsets[N] = len(edges)

with open(OUTPUT, "wb") as out:
    out.write(struct.pack("<I", N))           
    out.write(struct.pack("<I", len(edges)))  
    out.write(offsets.tobytes())
    out.write(neighbors.tobytes())
    out.write(weights.tobytes())

print("Wrote:", OUTPUT)