package main

import (
	//"fmt"
)

//essentially sweep might be a kind of finite state machine (fsm)

var IN_STATE [STATE_SIZE]int
var OUT_STATE [STATE_SIZE]int

func think(ir int, sensor_data_string string) int {
//	fmt.Println("DATA IN: ",sensor_data_string)
	var ACCUMULATORS [3]int
	var sensor_data []int

	sensor_data = convert_sensor_data(sensor_data_string)

	//var data_in = [0,0,0];
	//data_in = payload;

	var ix int
	var max_index int
	var max_value int

//	for i := 0; i < len(ACCUMULATORS); i++ {
//		ACCUMULATORS[i] = 0
//	}

	sweep(ir, sensor_data)

	/*possibles := (STATE_SIZE - INPS_SIZE)
	modo := possibles / INPS_SIZE
	fmt.Println("POSSIBLES,MODO: ",possibles,modo)

	for ak := 0; ak < possibles; ak++ {
		ix = ak % modo
		iak := 0
		iak = INPS_SIZE + ak
		fmt.Println("IX: ",ix," IAK: ",iak)
		ACCUMULATORS[ix] = ACCUMULATORS[ix] + OUT_STATE[iak]
	}
	*/
	//fmt.Println("OUTSTATE: ",OUT_STATE)
	step := len(ACCUMULATORS) //3
	for j:=0;j<STATE_SIZE;j++ {
		ix = j % step
		ACCUMULATORS[ix] += OUT_STATE[j]
	}
	//fmt.Println("ACC: ", ACCUMULATORS)

	for jj := 0; jj < len(ACCUMULATORS); jj++ {
		if ACCUMULATORS[jj] > max_value {
			max_value = ACCUMULATORS[jj]
			max_index = jj
		}
	}
	//if you don't know what you are doing, go straight
	if ACCUMULATORS[0] == ACCUMULATORS[1] && ACCUMULATORS[1] == ACCUMULATORS[2] {
		max_index = 1
	}

	return max_index
}

func sweep(ir int, data_in []int) {
	//gate_type := 0
	input1 := 0
	input2 := 0
	out := 0

	//might have to deep copy this
	for j := 0; j < STATE_SIZE; j++ {
		IN_STATE[j] = OUT_STATE[j]
	}

	//write over input section
	for j := 0; j < len(data_in); j++ {
		IN_STATE[j] = data_in[j]
	}
	for j := 0; j < STATE_SIZE; j++ {
		OUT_STATE[j] = 0
	}

	//var LUTS [STATE_SIZE][8]int
	var LUTS [NUM_NEURONS][8]int
	//I think this is just a copy...maybe..
	LUTS = rovers[ir].Luts
	var tt [2][2]int
	//big loop
	for ni := 0; ni < len(LUTS); ni++ {
		input1 = IN_STATE[LUTS[ni][0]]
		input2 = IN_STATE[LUTS[ni][1]]
		out = -9
		//ok here's where the metal hits the road
		//make a tt or figure out indexing ?
		//fine.. do the loop for now.
		//Put the tt in the struct ?
		knt := 4
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				tt[i][j] = LUTS[ni][knt]
				knt++
			}
		}
		out = tt[input1][input2]

		//only update outstate if out is 1
		//Only overwrite 0
		if out == 1 {
			OUT_STATE[LUTS[ni][2]] = out
			OUT_STATE[LUTS[ni][3]] = out
		}
	} //end of loop on NEURONS
	//end of sweep
}

func convert_sensor_data(sensor_data_string string) []int {
	//could have called Atoi on this, but meh
	var sig int
	sig = 0
	var sensor_data []int

	for i := 0; i < len(sensor_data_string); i++ {
		if sensor_data_string[i] == '0' {
			sig = 0
		} else {
			sig = 1
		}
		sensor_data = append(sensor_data,sig)
	}

	//throw in some noise if the sensors sense nothing...Just vast
	//empty space....
	knt := 0
	for ik := 0; ik < len(sensor_data); ik++ {
		knt = knt + int(sensor_data[ik])
	}

	if knt <= 0 {
		sensor_data[getRandomInt(0, len(sensor_data))] = 1
	}

	return sensor_data
}
