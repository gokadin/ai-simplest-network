# ai-linear-associative-network

This is the simplest neural network possible. 

## Overview

### What is a perceptron?

It's the simplest representation of a neuron. 
In essence, a perceptron has one or multiple inputs <img src="/tex/9fc20fb1d3825674c6a279cb0d5ca636.svg?invert_in_darkmode&sanitize=true" align=middle width=14.045887349999989pt height=14.15524440000002pt/> each having a specific weight <img src="/tex/c2a29561d89e139b3c7bffe51570c3ce.svg?invert_in_darkmode&sanitize=true" align=middle width=16.41940739999999pt height=14.15524440000002pt/> and one output <img src="/tex/deceeaf6940a8c7a5a02373728002b0f.svg?invert_in_darkmode&sanitize=true" align=middle width=8.649225749999989pt height=14.15524440000002pt/>. 

![alt text](readme-images/perceptron.jpg)

At the simplest level, the perceptron's output is the sum of it's inputs times its weights. 
<p align="center"><img src="/tex/c2d2775d67e954682fac686e557baed2.svg?invert_in_darkmode&sanitize=true" align=middle width=88.33802834999999pt height=44.89738935pt/></p>

### A simple example

![alt text](readme-images/perceptron-example.jpg)

Say we have a perceptron with two inputs <img src="/tex/f9b6dcc9279f659321ac3e1098b0ba4f.svg?invert_in_darkmode&sanitize=true" align=middle width=59.69172164999999pt height=21.18721440000001pt/> and <img src="/tex/bf84a893effff44b6d014b2b60460585.svg?invert_in_darkmode&sanitize=true" align=middle width=59.69172164999999pt height=21.18721440000001pt/>, with weights <img src="/tex/a4d0b18eae0e483561b7109b3e60efab.svg?invert_in_darkmode&sanitize=true" align=middle width=62.06524169999999pt height=21.18721440000001pt/> and <img src="/tex/b069935a7e85d35a3a6cfdb368977f8e.svg?invert_in_darkmode&sanitize=true" align=middle width=62.06524169999999pt height=21.18721440000001pt/>.  
Then the output <p align="center"><img src="/tex/8dfa08d909b122145492276ec756f3fa.svg?invert_in_darkmode&sanitize=true" align=middle width=8.64922575pt height=10.2739725pt/></p> will be
<p align="center"><img src="/tex/48c4f6073c4655b74cebf396493c9228.svg?invert_in_darkmode&sanitize=true" align=middle width=322.4824614pt height=13.789957499999998pt/></p>

## Activation function

Normally a perceptron processes its inputs through an activation function in order to normalise it's output. 

Examples of activation functions are:
- ReLU or rectified linear unit

<p align="center"><img src="/tex/3959b939e17b3c21d8589ea174a0fbfa.svg?invert_in_darkmode&sanitize=true" align=middle width=124.13822189999999pt height=16.438356pt/></p>

- Sigmoid

<p align="center"><img src="/tex/05ea2a27ecacfef18e873fcdafa3c70f.svg?invert_in_darkmode&sanitize=true" align=middle width=110.4028497pt height=34.3600389pt/></p>

- tanh

<p align="center"><img src="/tex/d7eb28cd142819e76d9f4c6df635ee71.svg?invert_in_darkmode&sanitize=true" align=middle width=110.05912829999998pt height=16.438356pt/></p>

## Learning