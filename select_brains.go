package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func select_brains() {
	sum := 0
	fmt.Println("\nFITNESS\n")

	for ir := 0; ir < NUM_ROVERS; ir++ {
		sum += rovers[ir].Fitness
	}
	fmt.Println("team sum score ", sum)

	//gotta sort an array with a slice sort
	sort.Slice(rovers[:], func(i, j int) bool {
		return rovers[i].Fitness > rovers[j].Fitness
	})
	fmt.Println("\n after sort")
	for ir := 0; ir < NUM_ROVERS; ir++ {
		fmt.Println(ir, rovers[ir].Fitness)
	}

	fmt.Println("\nBEST SCORE ", rovers[0].Fitness)
	fmt.Println("WRST SCORE ", rovers[NUM_ROVERS-1].Fitness)
	//zero out all the scores -- starting another epoch
	for ir := 0; ir < NUM_ROVERS; ir++ {
		rovers[ir].Fitness = 0
		rovers[ir].Dead = false
		rovers[ir].Xpos = arena.Width / 2
		rovers[ir].Ypos = arena.Height / 2
		rovers[ir].Vel_x = getRandomInt(-1, 2)
		rovers[ir].Vel_y = getRandomInt(-1, 2)
		rovers[ir].Angle_index = getRandomInt(0,8)
	}

	//HOPE ! HAHAHaaa
	//I hope this works. If lut is an array, than no problem
	//If lut is a slice, got problems. Because the lut copy
	//is the same as the lut it was copied from (*pointer stuff)
	//go doesn't have a deep copy function

	elite_cut := int(float64(NUM_ROVERS) * .2)
	for ib := elite_cut; ib < NUM_ROVERS; ib++ {
		bam := getRandomInt(0, elite_cut)
		c := [NUM_NEURONS][8]int{}
		c = rovers[bam].Luts
		rovers[ib].Luts = c
	}

	//fmt.Println("ROVERS: ",rovers)

	mutate_brains(elite_cut)

} //end of select

func getRandomFloat64(min float64, max float64) float64 {
	return 0.0 + (rand.Float64() * (max - min)) + min
}

func getRandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

func mutate_brains(elite_cut int) {
	//I am not mutating sign here. Too drastic
	//fmt.Println("IN MUTATE BRAINS")
	var num_mutations int
	var nn float64
	nn = float64(NUM_NEURONS)
	//num_mutations = int(nn * nn / 5.0)
	num_mutations = int( nn / 5.0)
	//fmt.Println("NUM MUTATIONS: ",num_mutations)
	for im := elite_cut; im < NUM_ROVERS; im++ {
		for k := 0; k < num_mutations; k++ {
			ix := getRandomInt(0, NUM_NEURONS)
			iy := getRandomInt(0, 8)
			if iy > 3 { //truth table 1/0
				if rovers[im].Luts[ix][iy] == 1 {
					rovers[im].Luts[ix][iy] = 0
				} else {
					rovers[im].Luts[ix][iy] = 1
				}
			}
			if iy <= 3 {
				rovers[im].Luts[ix][iy] = getRandomInt(0, STATE_SIZE)
			}

		}
	} //end of loop on num_rovers
} //end of mutate func
