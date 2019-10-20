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
		weight: rand.Float64(),
	}
}
