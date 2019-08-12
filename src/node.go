package main

type node struct {
	outputs []*connection
	value float64
	Delta float64
}

func newNode() *node {
	return &node{}
}

func (n *node) add(value float64) {
	n.value += value
}
