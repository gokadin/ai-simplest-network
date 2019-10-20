package main

import (
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
	
	inputLayer.attachTo(outputLayer)
	
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
