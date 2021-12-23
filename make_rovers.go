package main

import (
//"time"
//"math/rand"
//"fmt"
)

func make_rovers() {
	var rover Rover
	for i := 0; i < NUM_ROVERS; i++ {
		rover.luts = make_luts()
		rover.Fitness = 0
		rover.Dead = false
		rover.Xpos = arena.Width/2 + getRandomInt(-10,10)
		rover.Ypos = arena.Height/2 + getRandomInt(-10,10)
		rover.Vel_x   =  0
		rover.Vel_y   =  0
		rovers[i] = rover
	} //end of for loop on num_rovers
} //end of make_rovers



func make_luts() [NUM_NEURONS][8] int {
	var luts [NUM_NEURONS][8]int

	for ilut:=0;ilut<NUM_NEURONS;ilut++{

    //inputs
    luts[ilut][0] = getRandomInt(0,STATE_SIZE)
    luts[ilut][1] = getRandomInt(0,STATE_SIZE)
    //outputs
    luts[ilut][2] = getRandomInt(0,STATE_SIZE)
    luts[ilut][3] = getRandomInt(0,STATE_SIZE)
    //truth table
    for tt:=4;tt<8;tt++ {
            luts[ilut][tt] = getRandomInt(0,2)
    }
    } //end of loop on ilut
return luts
}
