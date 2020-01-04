# Simplest artificial neural network

This is the simplest artificial neural network possible explained and demonstrated. 

Try it online: [![Run on Repl.it](https://repl.it/badge/github/gokadin/ai-simplest-network)](https://repl.it/github/gokadin/ai-simplest-network)

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
The simplest form of an artificial neuron has one or multiple inputs <img src="/tex/9fc20fb1d3825674c6a279cb0d5ca636.svg?invert_in_darkmode&sanitize=true" align=middle width=14.045887349999989pt height=14.15524440000002pt/> each having a specific weight <img src="/tex/c2a29561d89e139b3c7bffe51570c3ce.svg?invert_in_darkmode&sanitize=true" align=middle width=16.41940739999999pt height=14.15524440000002pt/> and one output <img src="/tex/deceeaf6940a8c7a5a02373728002b0f.svg?invert_in_darkmode&sanitize=true" align=middle width=8.649225749999989pt height=14.15524440000002pt/>. 

![alt text](readme-images/perceptron.jpg)

At the simplest level, the output is the sum of its inputs times its weights. 
<p align="center"><img src="/tex/c2d2775d67e954682fac686e557baed2.svg?invert_in_darkmode&sanitize=true" align=middle width=88.33802834999999pt height=44.89738935pt/></p>

### A simple example

Say we have a network with two inputs <img src="/tex/f9b6dcc9279f659321ac3e1098b0ba4f.svg?invert_in_darkmode&sanitize=true" align=middle width=59.69172164999999pt height=21.18721440000001pt/> and <img src="/tex/bf84a893effff44b6d014b2b60460585.svg?invert_in_darkmode&sanitize=true" align=middle width=59.69172164999999pt height=21.18721440000001pt/> and two weights <img src="/tex/4b4518f1b7f0fb1347fa21506ebafb19.svg?invert_in_darkmode&sanitize=true" align=middle width=18.32105549999999pt height=14.15524440000002pt/> and <img src="/tex/f7eb0e840408d84a0c156d6efb611f3e.svg?invert_in_darkmode&sanitize=true" align=middle width=18.32105549999999pt height=14.15524440000002pt/>.  

The idea is to adjust the weights in such a way that the given inputs produce the desired output. 

Weights are normally initialized randomly since we can't know their optimal value ahead of time, however for simplicity we will initialize them both with <img src="/tex/f58ed17486d1735419372f2b7d091779.svg?invert_in_darkmode&sanitize=true" align=middle width=21.00464354999999pt height=21.18721440000001pt/>. 

![alt text](readme-images/perceptron-example.jpg)

Then the output <img src="/tex/deceeaf6940a8c7a5a02373728002b0f.svg?invert_in_darkmode&sanitize=true" align=middle width=8.649225749999989pt height=14.15524440000002pt/> will be
<p align="center"><img src="/tex/48c4f6073c4655b74cebf396493c9228.svg?invert_in_darkmode&sanitize=true" align=middle width=322.4824614pt height=13.789957499999998pt/></p>

### The error

If the output <img src="/tex/deceeaf6940a8c7a5a02373728002b0f.svg?invert_in_darkmode&sanitize=true" align=middle width=8.649225749999989pt height=14.15524440000002pt/> doesn't match the expected result, then we have an error.  
For example, if we wanted to get an expected output of <img src="/tex/ad35a4143e0a34d97d3abc63c4dc81a3.svg?invert_in_darkmode&sanitize=true" align=middle width=56.092022249999985pt height=21.18721440000001pt/> then we would have a difference of 

<p align="center"><img src="/tex/c744817f1f470ba09c3750aadef1c2a9.svg?invert_in_darkmode&sanitize=true" align=middle width=168.85052579999999pt height=13.789957499999998pt/></p>

The most common way to measure the error is to use the square difference:

<p align="center"><img src="/tex/215d8df8edc921e2d5c6c45e3cf05508.svg?invert_in_darkmode&sanitize=true" align=middle width=108.41233259999998pt height=32.990165999999995pt/></p>

If we had multiple associations of inputs and expected outputs, then the error becomes the sum of each association. 

<p align="center"><img src="/tex/56def9cd4b32408db97f6999c6ba45ca.svg?invert_in_darkmode&sanitize=true" align=middle width=154.52925015pt height=44.89738935pt/></p>

To rectify the error, we would need to adjust the weights in a way that the actual output matches the expected output. In our example, lowering <img src="/tex/4b4518f1b7f0fb1347fa21506ebafb19.svg?invert_in_darkmode&sanitize=true" align=middle width=18.32105549999999pt height=14.15524440000002pt/> from <img src="/tex/f58ed17486d1735419372f2b7d091779.svg?invert_in_darkmode&sanitize=true" align=middle width=21.00464354999999pt height=21.18721440000001pt/> to <img src="/tex/cde2d598001a947a6afd044a43d15629.svg?invert_in_darkmode&sanitize=true" align=middle width=21.00464354999999pt height=21.18721440000001pt/> would do the trick, since 
<p align="center"><img src="/tex/e6f831d1a270623d0d7f7ed67ad50360.svg?invert_in_darkmode&sanitize=true" align=middle width=243.73618499999998pt height=13.789957499999998pt/></p>

However, in order to adjust the weights of our neural networks for many different inputs and expected outputs, we need a *learning algorithm*. 

### Gradient descent

The idea is to use the error in order to adjust each weight so that the error is minimized.  

##### What is a gradient?

It's essentially a vector pointing to the direction of the steepest ascent of a function. The gradient is denoted with <img src="/tex/47c28f1929c18f887420345e9225e08b.svg?invert_in_darkmode&sanitize=true" align=middle width=13.69867124999999pt height=22.465723500000017pt/> and is simply the partial derivative of each variable of a function expressed as a vector.  

Example for a two variable function:

<p align="center"><img src="/tex/b142e84f3f77e6dc3144eb723cd4510d.svg?invert_in_darkmode&sanitize=true" align=middle width=303.75993285pt height=37.9216761pt/></p>

##### What is gradient descent?

The *descent* part simply means using the gradient to find the direction of steepest ascent of our function and then going in the opposite direction by a *small* amount many times to find the function *global minimum*.  

We use a constant called the **learning rate**, denoted with <img src="/tex/7ccca27b5ccc533a2dd72dc6fa28ed84.svg?invert_in_darkmode&sanitize=true" align=middle width=6.672392099999992pt height=14.15524440000002pt/> to define how small of a step to take in that direction.  

If <img src="/tex/7ccca27b5ccc533a2dd72dc6fa28ed84.svg?invert_in_darkmode&sanitize=true" align=middle width=6.672392099999992pt height=14.15524440000002pt/> is too large, then we risk overshooting the function minimum, but if it's too low then the network will take longer to learn and we risk getting stuck in a local minimum. 

![alt text](readme-images/gradient-descent.jpg)

##### Gradient descent applied to our example network

For our two weights <img src="/tex/4b4518f1b7f0fb1347fa21506ebafb19.svg?invert_in_darkmode&sanitize=true" align=middle width=18.32105549999999pt height=14.15524440000002pt/> and <img src="/tex/f7eb0e840408d84a0c156d6efb611f3e.svg?invert_in_darkmode&sanitize=true" align=middle width=18.32105549999999pt height=14.15524440000002pt/> we need to find the gradient of those weights with respect to the error function <img src="/tex/84df98c65d88c6adf15d4645ffa25e47.svg?invert_in_darkmode&sanitize=true" align=middle width=13.08219659999999pt height=22.465723500000017pt/>  

<p align="center"><img src="/tex/912be46ac0db99c8544f0800527d4b9f.svg?invert_in_darkmode&sanitize=true" align=middle width=147.62782815pt height=36.2778141pt/></p>

For both <img src="/tex/4b4518f1b7f0fb1347fa21506ebafb19.svg?invert_in_darkmode&sanitize=true" align=middle width=18.32105549999999pt height=14.15524440000002pt/> and <img src="/tex/f7eb0e840408d84a0c156d6efb611f3e.svg?invert_in_darkmode&sanitize=true" align=middle width=18.32105549999999pt height=14.15524440000002pt/>, we can find the gradient by using the chain rule

<p align="center"><img src="/tex/94bed65fa8ab5f8d63836d674b61da83.svg?invert_in_darkmode&sanitize=true" align=middle width=599.3444836499999pt height=39.452455349999994pt/></p>

From now on we will denote the <img src="/tex/d0a8416652c9aafee9f239948c2cd2e4.svg?invert_in_darkmode&sanitize=true" align=middle width=83.81122859999999pt height=28.92634470000001pt/> as the <img src="/tex/38f1e2a089e53d5c990a82f284948953.svg?invert_in_darkmode&sanitize=true" align=middle width=7.928075099999989pt height=22.831056599999986pt/> term. 

Once we have the gradient, we can update our weights

<p align="center"><img src="/tex/f46613c78403dce8eed0b6093ff36d28.svg?invert_in_darkmode&sanitize=true" align=middle width=222.19498950000002pt height=15.52509255pt/></p>

<p align="center"><img src="/tex/223d39de0113b6136f26962ed907c8aa.svg?invert_in_darkmode&sanitize=true" align=middle width=222.19498950000002pt height=15.52509255pt/></p>

And we repeat this process until the error is approximately 0​. 

## Code example

The included example teaches the following dataset to a neural network with two inputs and one output using gradient descent:

<p align="center"><img src="/tex/0cdd43e831c22b1560861b7a3e660010.svg?invert_in_darkmode&sanitize=true" align=middle width=233.52364695pt height=39.452455349999994pt/></p>

Once learned, the network should output ~0​ when given two <img src="/tex/034d0a6be0424bffe9a6e7ac9236c0f5.svg?invert_in_darkmode&sanitize=true" align=middle width=8.219209349999991pt height=21.18721440000001pt/>s and ~<img src="/tex/034d0a6be0424bffe9a6e7ac9236c0f5.svg?invert_in_darkmode&sanitize=true" align=middle width=8.219209349999991pt height=21.18721440000001pt/> when given a <img src="/tex/034d0a6be0424bffe9a6e7ac9236c0f5.svg?invert_in_darkmode&sanitize=true" align=middle width=8.219209349999991pt height=21.18721440000001pt/> and a <img src="/tex/29632a9bf827ce0200454dd32fc3be82.svg?invert_in_darkmode&sanitize=true" align=middle width=8.219209349999991pt height=21.18721440000001pt/>. 

## References

- Artificial intelligence engines by James V Stone (2019)