package main

type node struct {
	connections []*connection
	value       float64
	delta       float64
}

func newNode() *node {
	return &node{}
}
