package main

import "math/rand"

type connection struct {
	node *node
	weight float64
	gradient float64
}

func newConnection(node *node) *connection {
	return &connection{
		node: node,
		weight: 0.1 + rand.Float64() * (0.9 - 0.1),
	}
}
