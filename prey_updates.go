package main
import (
//	"fmt"
)

func do_prey_updates() [][3]int {
	var prey_positions [][3] int
	var junk  [3]int
	for i := 0; i < len(prey); i++ {
		if prey[i].Dead == 0 {
			junk[0] = prey[i].Xpos
			junk[1] = prey[i].Ypos
			junk[2] = prey[i].Dead
			prey_positions = append(prey_positions,junk)
		}

	} //end of loop on prey 
	//fmt.Println("prey positions: ",prey_positions)
	return prey_positions
} //end of do_update

func reset_prey() {
	for i := 0; i < len(prey); i++ {
                prey[i].Dead = 0
	}
}
func num_dead_prey() int {
	knt := 0
	for i := 0; i < len(prey); i++ {
                knt += prey[i].Dead 
	}
	return knt
}

