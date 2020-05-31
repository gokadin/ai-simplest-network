# Simplest artificial neural network

This is the simplest artificial neural network possible explained and demonstrated. 

## This is part 1 of a series of github repos on neural networks

- part 1 - simplest network (**you are here**)
- [part 2 - backpropagation](https://github.com/gokadin/ai-backpropagation)
- [part 3 - backpropagation-continued](https://github.com/gokadin/ai-backpropagation-continued)

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

The purpose of a network is to learn a certain output $y$ given certain input(s) $x$ by approximating a complex function with many parameters $w$ that we couldn't come up with ourselves. 

Say we have a network with two inputs $x_1 = 0.2$ and $x_2 = 0.4$ and two weights $w_1$ and $w_2$.  

The idea is to adjust the weights in such a way that the given inputs produce the desired output. 

Weights are normally initialized randomly since we can't know their optimal value ahead of time, however for simplicity we will initialize them both to $1$. 

![alt text](readme-images/perceptron-example.jpg)

If we calculate the output of this network, we will get 
$$ y = x_1 w_1 + x_2 w_2 = 0.2 * 1.0 + 0.4 * 1.0 = 0.6$$

### The error

If the output $y$ doesn't match the expected target value, then we have an error.  
For example, if we wanted to get a target value of $t = 0.5$ then we would have a difference of 

$$ y - t = 0.6 - 0.5 = 0.1$$

One common way to measure the error (also referred to as the cost function) is to use the mean squared error:

$$ E = \frac{1}{2} (y - t)^2 $$

If we had multiple associations of inputs and expected outputs, then the error becomes the sum of each association. 

$$ E = \frac{1}{2n} \sum_{i=1}^n (y_i - t_i)^2 $$

We use the mean squared error to measure how far away the results are from our desired target. The squaring removes negative signs and gives more weight to bigger differences in result. 

To rectify the error, we would need to adjust the weights in a way that the output matches our target. In our example, lowering $w_1$ from $1.0$ to $0.5$ would do the trick, since 
$$ y = t = 0.2 * 0.5 + 0.4 * 1.0 = 0.5 $$

However, in order to adjust the weights of our neural networks for many different inputs and target values, we need a *learning algorithm* to do this for us automatically. 

### Gradient descent

The idea is to use the error to infer how each weight should be adjusted so that the error is minimized, but first, we need to learn about gradients. 

##### What is a gradient?

It's essentially a vector pointing to the direction of the steepest ascent of a function. The gradient is denoted with $\nabla$ and is simply the partial derivative of each variable of a function expressed as a vector. 

It looks like this for a two variable function:

$$ \nabla f(x, y) = \langle f_x, f_y \rangle = \bigl \langle \frac{\partial f(x, y)}{\partial x}, \frac{\partial f(x, y)}{\partial y} \bigl \rangle $$

Let's inject some numbers and calculate the gradient with a simple example. 
Say we have a function $f(x,y) = 2x^2 + 3y^3$, then the gradient would be

$$\nabla f(x,y) = \bigl \langle4x, 9y^2 \bigl \rangle$$

##### What is gradient descent?

The *descent* part simply means using the gradient to find the direction of steepest ascent of our function and then going in the opposite direction by a *small* amount many times to find the function *global (or sometimes local) minimum*.  

We use a constant called the **learning rate**, denoted with $\epsilon$ to define how small of a step to take in that direction.  

If $\epsilon$ is too large, then we risk overshooting the function minimum, but if it's too low then the network will take longer to learn and we risk getting stuck in a shallow local minimum. 

![alt text](readme-images/gradient-descent.jpg)

##### Gradient descent applied to our example network

For our two weights $w_1$ and $w_2$ we need to find the gradient of those weights with respect to the error function $E$  

$$ \nabla_w E = \bigl \langle \frac{\partial E}{\partial w_1}, \frac{\partial E}{\partial w_2} \bigl \rangle $$

For both $w_1$ and $w_2$, we can find the gradient by using the chain rule

$$ \frac{\partial E}{\partial w_i} = \frac{\partial E}{\partial y}\frac{\partial y}{\partial w_i} = \frac{\partial}{\partial y}\left(\frac{1}{2}(y - y\prime)^2\right) \frac{\partial}{\partial w_i}\left(x_iw_i\right) = (y - t) \frac{\partial}{\partial y}\left(y - t\right) x_i = (y - t)x_i $$

From now on we will denote the $\frac{\partial E}{\partial y} = y - t$ as the $\delta$ term for simplicity. 

Once we have the gradient, we can update our weights

$$ w_1 = w_1 - \epsilon \nabla_{w_1}E = w_1 -\epsilon \delta x_1 $$

$$ w_2 = w_2 - \epsilon \nabla_{w_2}E = w_2 -\epsilon \delta x_2 $$

And we repeat this process until the error is minimized and is close enough to zero for our comfort. 

## Code example

The included example teaches the following dataset to a neural network with two inputs and one output using gradient descent:

$$ x = \begin{bmatrix}
    1.0 & 1.0 \\
    1.0 & 0.0 \\
\end{bmatrix} \quad y\prime = \begin{bmatrix}
    0.0 & 1.0 \\
\end{bmatrix} $$

Once learned, the network should output ~0 when given two $1$s and ~$1$ when given a $1$ and a $0$. 

### How to run

#### Online on repl.it

[![Run on Repl.it](https://repl.it/badge/github/gokadin/ai-simplest-network)](https://repl.it/github/gokadin/ai-simplest-network)

#### Docker

``` bash
docker build -t simplest-network .
docker run --rm simplest-network
```

## References

1. Artificial intelligence engines by James V Stone (2019)
2. Complete guide on deep learning: http://neuralnetworksanddeeplearning.com/chap2.html