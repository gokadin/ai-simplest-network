package main

import (
	"log"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	inputLayer := NewLayer(2)
	outputLayer := NewLayer(1)

	inputLayer.AttachTo(outputLayer)

	inputs := make([][]float64, 0)
    inputs = append(inputs, []float64{1.0, 1.0})
	inputs = append(inputs, []float64{1.0, 0.0})
	outputs := make([][]float64, 0)
	outputs = append(outputs, []float64{0.0})
	outputs = append(outputs, []float64{1.0})

	Learn(inputs, outputs, inputLayer, outputLayer)
	log.Println("Results ----------------------")
	Run(inputs, inputLayer, outputLayer)
}

func Run(inputs [][]float64, inputLayer *Layer, outputLayer *Layer) {
	for _, input := range inputs {
		inputLayer.ResetValues()

		for i, node := range inputLayer.nodes {
			node.Add(input[i])
		}

		inputLayer.Activate()

		log.Println("Inputs", input[0], input[1], "produced", outputLayer.nodes[0].value)
	}
}

func Learn(inputs [][]float64, outputs [][]float64, inputLayer *Layer, outputLayer *Layer) {
	counter := 0
	learningRate := 0.1

	learn := true
	for learn {
		inputLayer.ResetDeltas()
		err := 0.0
		for inputIndex, input := range inputs {
			counter++

			// run network

			inputLayer.ResetValues()
			for i, node := range inputLayer.nodes {
				node.Add(input[i])
			}
			inputLayer.Activate()

			// calculate error

			delta := outputLayer.nodes[0].value - outputs[inputIndex][0]
			err += math.Pow(delta, 2)

			for _, inputNode := range inputLayer.nodes {
				derivative := delta * inputNode.value
				deltaWeight := -learningRate * derivative
				inputNode.Delta += deltaWeight
			}
		}

		err /= 2
		log.Println("err: ", err)
		if err < 0.001 {
			learn = false
		}

		// update weights

		for _, inputNode := range inputLayer.nodes {
			inputNode.outputs[0].weight += inputNode.Delta
		}
	}

    log.Println("Finished after", counter, "iterations")
}

type Node struct {
	outputs []*Connection
	value float64
	Delta float64
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) Add(value float64) {
	n.value += value
}

type Layer struct {
	nodes []*Node
	nextLayer *Layer
}

func NewLayer(size int) *Layer {
	nodes := make([]*Node, size)
	for i := range nodes {
        nodes[i] = NewNode()
	}
	return &Layer{
		nodes: nodes,
	}
}

func (l *Layer) AttachTo(nextLayer *Layer) {
	l.nextLayer = nextLayer

	for _, node := range l.nodes {
		for _, nextNode := range nextLayer.nodes {
			node.outputs = append(node.outputs, NewConnection(nextNode))
		}
	}
}

func (l *Layer) Activate() {
    for _, node := range l.nodes {
		for _, connection := range node.outputs {
			connection.node.Add(node.value * connection.weight)
		}
	}

    if l.nextLayer != nil {
		l.nextLayer.Activate()
	}
}

func (l *Layer) ResetValues() {
	for _, node := range l.nodes {
		node.value = 0.0
	}

	if l.nextLayer != nil {
		l.nextLayer.ResetValues()
	}
}

func (l *Layer) ResetDeltas() {
	for _, node := range l.nodes {
		node.Delta = 0.0
	}

	if l.nextLayer != nil {
		l.nextLayer.ResetDeltas()
	}
}

type Connection struct {
	node *Node
	weight float64
}

func NewConnection(node *Node) *Connection {
	return &Connection{
		node: node,
		weight: 0.1 + rand.Float64() * (0.9 - 0.1),
	}
}