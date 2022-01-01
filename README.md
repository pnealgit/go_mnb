# go_mnb
Markov Network Brains with GA in Go

Written by Phillip R. Neal
December 23, 2021

December 24, 2021 - Got it working. The GO->HTML5 model goes too fast.
Which is a good thing. But hard to watch stuff happening. 


December 26, 2021 - Working with Filled rectangles as prey. 

December 29, 2021 - Start to work on only showing 1 rover eating grid
of food items

December 30, 2021 - When I go down to one rover on the html5 page
I start to lose behavior. Antenna don't make it all the way, etc.
Could be FPS. I will try to send an ack from the browser back to the
server.

January 1, 2022 - Implemented the ack. Made screen update a function of
data received, not just an interval.


Markov Simple Markov Network Brain with Go talking to
an HTML/JS graphic page vi websockets

PLEASE NOTE: 

This is not a "framework". This is a Go/HTML5 program.

Background links:
My own take on networked brains ala Chris Adami's and
his teams ideas.

See: 

 Chris Adami's site (http://adamilab.msu.edu/markov-network-brains/)

 Another good description (http://devosoft.org/a-quick-introduction-to-markov-network-brains/)

 Jefferey Cave's really neat implementation
 (https://medium.com/@jefferey.cave/fun-with-markov-network-brains-8041c35ca883)

 Arend Hintze's "Markov Brains: A Technical Introduction" is also great. (https://arxiv.org/pdf/1709.05601v1.pdf)

An implementation in Python at:
https://github.com/nicholasharris/Markov-Brains-Python/blob/master/markov.py

Evolution:

1. Generate a population of "BRAINS".
2. Evaluate the fitness of each of the original BRAINS
3. Sort the population descending based on fitness
4. With the best BRAINS in population after sort)
    - Replace and mutate all other brains depending on mutation rate

5. Evaluate the fitness of all the new brains 
6. If candidate brain is better than the worst BRAIN, 
   replace worst Brain with candidate brain
7. Go to 4 above.

The Goal:

1. Have all the rovers head for food 
3. The rover accumulates fitness depending on how many times in a row
   it goes straight as well as if it touches food
 

How to Run:
1. Download the repository one way or another.
2. At the CLI type "./run.sh . This will fire up the GO side
3. Open a Chrome browser (maybe others will work)
3. Point it at "localhost:8081"
4. You should see a bunch of green static food circles and
   a bunch of moving colored circles.

Overview:

1. Rovers have 3 sensors.
2. The sensors detect whether or not it senses a wall or the side of a cube.
and the distance to said detection
3. The data returned from the sensors is sent to the think module
4. The think module contains the Markov Network Brain (MNB).
5. The think module returns the best sensor choice to use.
6. The rover turns in the direction of the best sensor.

Gates:

1. I use a simple gate structure.
    - 2 input indexes 
    - 2 output indexes
    - a 2x2 matrix of 0 or 1 randomly generated for the truth table

2. I do not use probabilistic, not,neuron,or learning gates.

3. The truth table for each gate is a 2d matrix indexed by the input values

4. The values for each input index points to the binary value at INPUT_STATE[input_index]

5. The values for each output index points to the OUT_STATE[output_index] where the result of the truth table will be written.


