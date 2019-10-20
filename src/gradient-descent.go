package main

import (
	"fmt"
	"math"
)

func learn(inputs [][]float64, outputs [][]float64, inputLayer *layer, outputLayer *layer) {
	counter := 0
	for {
		counter++
		err := 0.0
		for inputIndex, input := range inputs {
			forwardPass(inputLayer, input)
			calculateDeltas(outputLayer, outputs[inputIndex])
			accumulateGradients(inputLayer)
			err += math.Pow(accumulateError(outputLayer, outputs[inputIndex]), 2)
		}

		err /= 2 * float64(len(outputLayer.nodes))
		fmt.Println("err: ", err)
		if err < 0.0001 {
			break
		}

		updateWeights(inputLayer)
	}

	fmt.Println("Finished after", counter, "iterations")
}

func calculateDeltas(outputLayer *layer, expectedOutput []float64) {
	for nodeIndex, outputNode := range outputLayer.nodes {
		outputNode.delta = outputNode.value - expectedOutput[nodeIndex]
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
