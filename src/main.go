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

func test(inputs [][]float64, inputLayer *layer, outputLayer *layer) {
	fmt.Println()
	fmt.Println("Results ----------------------")
	
	for _, input := range inputs {
		forwardPass(inputLayer, input)

		outputs := make([]float64, len(outputLayer.nodes))
		for outputNodeIndex, outputNode := range outputLayer.nodes {
            outputs[outputNodeIndex] = outputNode.value
		}
        fmt.Println(input, "=>", outputs)
	}
}

func learn(inputs [][]float64, outputs [][]float64, inputLayer *layer, outputLayer *layer) {
	counter := 0
	for {
		inputLayer.resetDeltas()
		err := 0.0
		for inputIndex, input := range inputs {
			counter++

			forwardPass(inputLayer, input)

            accumulateDeltas(outputLayer, outputs[inputIndex])

			accumulateGradients(inputLayer)

			err += math.Pow(accumulateError(outputLayer, outputs[inputIndex]), 2)
		}

		err /= 2
		fmt.Println("err: ", err)
		if err < 0.0001 {
            break
		}

		updateWeights(inputLayer)
	}

    fmt.Println("Finished after", counter, "iterations")
}

func accumulateDeltas(outputLayer *layer, expectedOutput []float64) {
	for nodeIndex, outputNode := range outputLayer.nodes {
		outputNode.delta += outputNode.value - expectedOutput[nodeIndex]
	}
}

func accumulateGradients(inputLayer *layer) {
	for _, inputNode := range inputLayer.nodes {
		for _, connection := range inputNode.connections {
            connection.gradient += connection.node.delta * inputNode.value
		}
	}
}

func accumulateError(outputLayer *layer, expectedOutput []float64) float64 {
	err := 0.0
	for nodeIndex, outputNode := range outputLayer.nodes {
        err += outputNode.value - expectedOutput[nodeIndex]
	}
	return err
}

func forwardPass(inputLayer *layer, input []float64) {
	inputLayer.resetValues()
	for i, node := range inputLayer.nodes {
		node.value += input[i]
	}
	inputLayer.activate()
}

func updateWeights(inputLayer *layer) {
	for _, inputNode := range inputLayer.nodes {
		for _, connection := range inputNode.connections {
			connection.weight -= learningRate * connection.gradient
			connection.gradient = 0.0
		}
	}
}
