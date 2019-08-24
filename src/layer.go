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

func (l *layer) attachTo(nextLayer *layer) {
	l.nextLayer = nextLayer

	for _, node := range l.nodes {
		for _, nextNode := range nextLayer.nodes {
			node.connections = append(node.connections, newConnection(nextNode))
		}
	}
}

func (l *layer) activate() {
	for _, node := range l.nodes {
		for _, connection := range node.connections {
			connection.node.value += node.value * connection.weight
		}
	}

	if l.nextLayer != nil {
		l.nextLayer.activate()
	}
}

func (l *layer) resetValues() {
	for _, node := range l.nodes {
		node.value = 0.0
	}

	if l.nextLayer != nil {
		l.nextLayer.resetValues()
	}
}

func (l *layer) resetDeltas() {
	for _, node := range l.nodes {
		node.delta = 0.0
	}

	if l.nextLayer != nil {
		l.nextLayer.resetDeltas()
	}
}

