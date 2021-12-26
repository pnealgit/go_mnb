package main

import (
//"time"
//"math/rand"
//"fmt"
)

func make_prey() {
	// box_width = box_height
	//went to filled rectangles
	for i := 0; i < NUM_PREY; i++ {
		prey[i].Xpos = getRandomInt(BOX_WIDTH, arena.Width-BOX_WIDTH)
		prey[i].Ulx = prey[i].Xpos - BOX_HALF
		prey[i].Ypos = getRandomInt(BOX_WIDTH, arena.Height-BOX_WIDTH)
		prey[i].Uly = prey[i].Ypos - BOX_HALF
		prey[i].Type = 1
		prey[i].Dead = false
	} //end of for loop on i
} //end of make_prey
