# ai-linear-associative-network

This is the simplest neural network possible. 

## Overview

### What is a perceptron?

It's the simplest representation of a neuron. 
In essence, a perceptron has one or multiple inputs $x_i$ each having a specific weight $w_i$ and one output $y$. 

![alt text](readme-images/perceptron.jpg)

At the simplest level, the perceptron's output is the sum of it's inputs times its weights. 
$$ y = \sum_{i=1}^n w_i x_i $$

### A simple example

![alt text](readme-images/perceptron-example.jpg)

Say we have a perceptron with two inputs $x_1 = 0.2$ and $x_2 = 0.4$, with weights $w_1 = 1.0$ and $w_2 = 1.0$.  
Then the output $$y$$ will be
$$ y = x_1 w_1 + x_2 w_2 = 0.2 * 1.0 + 0.4 * 1.0 = 0.6$$

## Activation function

Normally a perceptron processes its inputs through an activation function in order to normalise it's output. 

Examples of activation functions are:
- ReLU or rectified linear unit

$$ f(x) = max(0, x) $$

- Sigmoid

$$ f(x) = \frac{1}{1 + e^{-x}} $$

- tanh

$$ f(x) = tanh(x) $$

## Learning