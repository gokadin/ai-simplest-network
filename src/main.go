package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const learningRate = 0.1

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	inputLayer, outputLayer := buildNetwork()
	inputs, outputs := generateData()

	learn(inputs, outputs, inputLayer, outputLayer)
	test(inputs, inputLayer, outputLayer)
}

func buildNetwork() (*layer, *layer) {
	inputLayer := newLayer(2)
	outputLayer := newLayer(1)
	
	inputLayer.AttachTo(outputLayer)
	
	return inputLayer, outputLayer
}

func generateData() ([][]float64, [][]float64) {
	inputs := make([][]float64, 0)
	inputs = append(inputs, []float64{1.0, 1.0})
	inputs = append(inputs, []float64{1.0, 0.0})
	
	outputs := make([][]float64, 0)
	outputs = append(outputs, []float64{0.0})
	outputs = append(outputs, []float64{1.0})
	
	return inputs, outputs
}

func test(inputs [][]float64, inputLayer *layer, outputLayer *layer) {
	fmt.Println()
	fmt.Println("Results ----------------------")
	
	for _, input := range inputs {
		forwardPass(inputLayer, input)
		fmt.Println("Inputs", input[0], input[1], "produced", outputLayer.nodes[0].value)
	}
}

func learn(inputs [][]float64, outputs [][]float64, inputLayer *layer, outputLayer *layer) {
	counter := 0
	learn := true
	for learn {
		inputLayer.ResetDeltas()
		err := 0.0
		for inputIndex, input := range inputs {
			counter++

			// run network
			forwardPass(inputLayer, input)

			// accumulate error
			delta := outputLayer.nodes[0].value - outputs[inputIndex][0]
			err += math.Pow(delta, 2)

			// accumulate gradient
			accumulateGradient(inputLayer, delta)
		}

		err /= 2
		fmt.Println("err: ", err)
		if err < 0.001 {
			learn = false
		}

		updateWeights(inputLayer)
	}

    fmt.Println("Finished after", counter, "iterations")
}

func accumulateGradient(inputLayer *layer, delta float64) {
	for _, inputNode := range inputLayer.nodes {
		derivative := delta * inputNode.value
		deltaWeight := derivative
		inputNode.Delta += deltaWeight
	}
}

func forwardPass(inputLayer *layer, input []float64) {
	inputLayer.ResetValues()
	for i, node := range inputLayer.nodes {
		node.add(input[i])
	}
	inputLayer.Activate()
}

func updateWeights(inputLayer *layer) {
	for _, inputNode := range inputLayer.nodes {
		inputNode.outputs[0].weight -= learningRate * inputNode.Delta
	}
}
