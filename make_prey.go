package main

import (
//"time"
//"math/rand"
//"fmt"
//"os"
//  "math"
)

func make_prey() {
	nrow := 3
	ncol := 5

	prey = nil

	deltax := (arena.Width-FOOD_RADIUS)/nrow
	deltay := (arena.Height-FOOD_RADIUS)/ncol
	//knt := 0
	var junk Prey
	for icol:= 0; icol < ncol; icol++ {
	for irow:= 0; irow < nrow; irow++ {

		junk.Xpos = irow * deltax + 2 * FOOD_RADIUS
		junk.Ypos = icol * deltay + 2 * FOOD_RADIUS
		junk.Dead = 0
		prey = append(prey,junk)
	} //end of for loop on irow
	} //end of for loop on icol
	//fmt.Println("PREY: ",prey)
	//os.Exit(22)
} //end of make_prey
