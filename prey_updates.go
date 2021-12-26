package main
import (
	//"fmt"
)

func do_prey_updates() [NUM_PREY][2]int {
	var prey_positions [NUM_PREY][2] int
	for i := 0; i < NUM_PREY; i++ {
		if prey[i].Dead {
                    prey[i].Xpos = getRandomInt(BOX_WIDTH, arena.Width-BOX_WIDTH)
                     prey[i].Ulx  = prey[i].Xpos - BOX_HALF
                     prey[i].Ypos = getRandomInt(BOX_WIDTH, arena.Height-BOX_WIDTH)
                     prey[i].Uly  = prey[i].Ypos - BOX_HALF
                }

		prey[i].Dead = false
		prey_positions[i][0] = prey[i].Ulx
		prey_positions[i][1] = prey[i].Uly
	} //end of loop on prey 
	return prey_positions
} //end of do_update

