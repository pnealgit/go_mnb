package main

import (
//"time"
//"math/rand"
//"fmt"
)

func make_rovers() {
	var rover Rover
	for i := 0; i < NUM_ROVERS; i++ {
		rover.Luts = make_Luts()
		rover.Fitness = 0
		rover.Dead = 0
		rover.Xpos = arena.Width/2 + getRandomInt(-10, 10)
		rover.Ypos = arena.Height/2 + getRandomInt(-10, 10)
		rover.Angle_index = getRandomInt(0,8)
		rover.Vel_x = 0
		rover.Vel_y = 0
		rover.Time_to_live = NUM_MAX_STEPS
		rovers[i] = rover
	} //end of for loop on num_rovers
} //end of make_rovers

func make_Luts() [NUM_NEURONS][8]int {
	var Luts [NUM_NEURONS][8]int

	for ilut := 0; ilut < NUM_NEURONS; ilut++ {

		//inputs
		Luts[ilut][0] = getRandomInt(0, STATE_SIZE)
		Luts[ilut][1] = getRandomInt(0, STATE_SIZE)
		//outputs
		Luts[ilut][2] = getRandomInt(0, STATE_SIZE)
		Luts[ilut][3] = getRandomInt(0, STATE_SIZE)
		//truth table
		for tt := 4; tt < 8; tt++ {
			Luts[ilut][tt] = getRandomInt(0, 2)
		}
	} //end of loop on ilut
	return Luts
}
