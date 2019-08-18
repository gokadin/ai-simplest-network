# Simplest artificial neural network

This is the simplest artificial neural network possible explained and demonstrated. 

## This is part 1 of a series of github repos on neural networks

- part 1 - simplest network (**you are here**)
- [part 2 - backpropagation](https://github.com/gokadin/ai-backpropagation)
- [part 3 - backpropagation-continued](https://github.com/gokadin/ai-backpropagation-continued)
- [part 4 - hopfield networks](https://github.com/gokadin/ai-hopfield-networks)

## Table of Contents

- [Theory](#theory)  
  - [Mimicking neurons](#mimicking-neurons)
  - [A simple example](#a-simple-example)
  - [The error](#the-error)
  - [Gradient descent](#gradient-descent)
- [Code example](#code-example)
- [References](#references)

## Theory

### Mimicking neurons

Artificial neural networks are inspired by the brain by having interconnected artificial neurons store patterns and communicate with each other. 
The simplest form of an artificial neuron has one or multiple inputs $x_i$ each having a specific weight $w_i$ and one output $y$. 

![alt text](readme-images/perceptron.jpg)

At the simplest level, the output is the sum of its inputs times its weights. 
$$ y = \sum_{i=1}^n w_i x_i $$

### A simple example

Say we have a network with two inputs $x_1 = 0.2$ and $x_2 = 0.4$ and two weights $w_1$ and $w_2$.  

The idea is to adjust the weights in such a way that the given inputs produce the desired output. 

Weights are normally initialized randomly since we can't know their optimal value ahead of time, however for simplicity we will initialize them both with $1.0$. 

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

### Gradient descent

The idea is to use the error in order to find how to adjust each weight so that the error is minimized.  

##### What is a gradient?

It's essentially a vector pointing to the direction of the steepest ascent of a function. The gradient is denoted with $\nabla$ and is simply the partial derivative of each variable of a function expressed as a vector.  

Example for a two variable function:

$$ \nabla f(x, y) = \langle f_x, f_y \rangle = \bigl \langle \frac{\partial f(x, y)}{\partial x}, \frac{\partial f(x, y)}{\partial y} \bigl \rangle $$

##### What is gradient descent?

The *descent* part simply means using the gradient to find the direction of steepest ascent of our function and then going in the opposite direction by a *small* amount many times to find the function *minimum*.  

We use a constant called the **learning rate**, denoted with $\epsilon$ to define how small of a step to take in that direction.  

If $\epsilon$ is too large, then we risk overshooting the function minimum. 

![alt text](readme-images/gradient-descent.jpg)

##### Gradient descent applied to our example network

For our two weights $w_1$ and $w_2$ we need to find the gradient of those weights with respect to the error function $E$  

$$ \frac{\partial E}{\partial w_1} = \delta x_1 \quad and \quad \frac{\partial E}{\partial w_2} = \delta x_2 $$

which we can write as a vector

$$ \nabla_w E = \bigl \langle \frac{\partial E}{\partial w_1}, \frac{\partial E}{\partial w_2} \bigl \rangle $$

Once we have the gradient, we can update our weights

$$ w = w - \epsilon \nabla E $$

And we repeat this process until the error is approximately $0$. 

## Code example

The included example teaches the following dataset to a neural network with two inputs and one output using gradient descent:

$$ x = \begin{bmatrix}
    1.0 & 1.0 \\
    1.0 & 0.0 \\
\end{bmatrix} \quad y\prime = \begin{bmatrix}
    0.0 & 1.0 \\
\end{bmatrix} $$

Once learned, the network should output ~$0$ when given two $1$s and ~$1$ when given a $1$ and a $0$. 

## References

- Artificial intelligence engines by James V Stone (2019)