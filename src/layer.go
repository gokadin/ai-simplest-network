package main

type layer struct {
	nodes []*node
	nextLayer *layer
}

func newLayer(size int) *layer {
	nodes := make([]*node, size)
	for i := range nodes {
		nodes[i] = newNode()
	}
	return &layer{
		nodes: nodes,
	}
}

func (l *layer) AttachTo(nextLayer *layer) {
	l.nextLayer = nextLayer

	for _, node := range l.nodes {
		for _, nextNode := range nextLayer.nodes {
			node.outputs = append(node.outputs, newConnection(nextNode))
		}
	}
}

func (l *layer) Activate() {
	for _, node := range l.nodes {
		for _, connection := range node.outputs {
			connection.node.add(node.value * connection.weight)
		}
	}

	if l.nextLayer != nil {
		l.nextLayer.Activate()
	}
}

func (l *layer) ResetValues() {
	for _, node := range l.nodes {
		node.value = 0.0
	}

	if l.nextLayer != nil {
		l.nextLayer.ResetValues()
	}
}

func (l *layer) ResetDeltas() {
	for _, node := range l.nodes {
		node.Delta = 0.0
	}

	if l.nextLayer != nil {
		l.nextLayer.ResetDeltas()
	}
}

