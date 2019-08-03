# Linear associative networks

This is the simplest artifical neural network possible. 

## Table of Contents
* [Overview](#overview)  
  * [What is a perceptron?](#what-is-a-perceptron)
  * [A simple example](#a-simple-example)
  * [The error](#the-error)

## Overview

### What is a perceptron?

It's the simplest conceptual representation of a neuron. 
In essence, a perceptron has one or multiple inputs $x_i$ each having a specific weight $w_i$ and one output $y$. 

![alt text](readme-images/perceptron.jpg)

At the simplest level, the perceptron's output is the sum of it's inputs times its weights. 
$$ y = \sum_{i=1}^n w_i x_i $$

### A simple example

Say we have a perceptron with two inputs $x_1 = 0.2$ and $x_2 = 0.4$, with weights $w_1 = 1.0$ and $w_2 = 1.0$.  

![alt text](readme-images/perceptron-example.jpg)

Then the output $y$ will be
$$ y = x_1 w_1 + x_2 w_2 = 0.2 * 1.0 + 0.4 * 1.0 = 0.6$$

### The error

If the output $y$ doesn't match the expected result, then we have an error.  
For example, if we wanted to get an expected output of $y\prime = 0.5$ then we would have a delta of 

$$ \delta = y - y\prime = 0.6 - 0.5 = 0.1$$

The most common way to measure the error is to use the square difference:

$$ E = \frac{1}{2} (y - y\prime)^2 = \frac{1}{2} \delta^2 $$

If we would have multiple sets of inputs and multiple sets of expected outputs, then the error becomes the sum of each set. 

$$ E = \frac{1}{2} \sum_{i=1}^n (y_i - y_i\prime)^2 = \frac{1}{2} \sum_{i=1}^n \delta_i^2 $$

To rectify the error, we would need to adjust the weights in a way that the actual output matches the expected output. In our example, lowering $w_1$ from $1.0$ to $0.5$ would do the trick, since 
$$ y = y\prime = 0.2 * 0.5 + 0.4 * 1.0 = 0.5 $$

However, in order to adjust the weights of our neural networks for many different inputs and expected outputs, we need a *learning algorithm*. 