# Linear associative networks

This is the simplest artifical neural network possible. 

## Overview

### What is a perceptron?

It's the simplest conceptual representation of a neuron. 
In essence, a perceptron has one or multiple inputs <img src="/tex/9fc20fb1d3825674c6a279cb0d5ca636.svg?invert_in_darkmode&sanitize=true" align=middle width=14.045887349999989pt height=14.15524440000002pt/> each having a specific weight <img src="/tex/c2a29561d89e139b3c7bffe51570c3ce.svg?invert_in_darkmode&sanitize=true" align=middle width=16.41940739999999pt height=14.15524440000002pt/> and one output <img src="/tex/deceeaf6940a8c7a5a02373728002b0f.svg?invert_in_darkmode&sanitize=true" align=middle width=8.649225749999989pt height=14.15524440000002pt/>. 

![alt text](readme-images/perceptron.jpg)

At the simplest level, the perceptron's output is the sum of it's inputs times its weights. 
<p align="center"><img src="/tex/c2d2775d67e954682fac686e557baed2.svg?invert_in_darkmode&sanitize=true" align=middle width=88.33802834999999pt height=44.89738935pt/></p>

### A simple example

Say we have a perceptron with two inputs <img src="/tex/f9b6dcc9279f659321ac3e1098b0ba4f.svg?invert_in_darkmode&sanitize=true" align=middle width=59.69172164999999pt height=21.18721440000001pt/> and <img src="/tex/bf84a893effff44b6d014b2b60460585.svg?invert_in_darkmode&sanitize=true" align=middle width=59.69172164999999pt height=21.18721440000001pt/>, with weights <img src="/tex/a4d0b18eae0e483561b7109b3e60efab.svg?invert_in_darkmode&sanitize=true" align=middle width=62.06524169999999pt height=21.18721440000001pt/> and <img src="/tex/b069935a7e85d35a3a6cfdb368977f8e.svg?invert_in_darkmode&sanitize=true" align=middle width=62.06524169999999pt height=21.18721440000001pt/>.  

![alt text](readme-images/perceptron-example.jpg)

Then the output <img src="/tex/deceeaf6940a8c7a5a02373728002b0f.svg?invert_in_darkmode&sanitize=true" align=middle width=8.649225749999989pt height=14.15524440000002pt/> will be
<p align="center"><img src="/tex/48c4f6073c4655b74cebf396493c9228.svg?invert_in_darkmode&sanitize=true" align=middle width=322.4824614pt height=13.789957499999998pt/></p>

### The error

If the output <img src="/tex/deceeaf6940a8c7a5a02373728002b0f.svg?invert_in_darkmode&sanitize=true" align=middle width=8.649225749999989pt height=14.15524440000002pt/> doesn't match the expected result, then we have an error.  
For example, if we wanted to get an expected output of <img src="/tex/ad35a4143e0a34d97d3abc63c4dc81a3.svg?invert_in_darkmode&sanitize=true" align=middle width=56.092022249999985pt height=21.18721440000001pt/> then we would have a delta of 

<p align="center"><img src="/tex/70587273e97df3ceb21ab1b1987c0c58.svg?invert_in_darkmode&sanitize=true" align=middle width=198.69622905pt height=14.611878599999999pt/></p>

and we would need to adjust the weights in order to rectify that error. In our example, lowering <img src="/tex/4b4518f1b7f0fb1347fa21506ebafb19.svg?invert_in_darkmode&sanitize=true" align=middle width=18.32105549999999pt height=14.15524440000002pt/> from <img src="/tex/f58ed17486d1735419372f2b7d091779.svg?invert_in_darkmode&sanitize=true" align=middle width=21.00464354999999pt height=21.18721440000001pt/> to <img src="/tex/cde2d598001a947a6afd044a43d15629.svg?invert_in_darkmode&sanitize=true" align=middle width=21.00464354999999pt height=21.18721440000001pt/> would do the trick, since 
<p align="center"><img src="/tex/e6f831d1a270623d0d7f7ed67ad50360.svg?invert_in_darkmode&sanitize=true" align=middle width=243.73618499999998pt height=13.789957499999998pt/></p>

However, in order to adjust the weights of our neural networks for many different inputs and expected outputs, we need a *learning algorithm*. 
