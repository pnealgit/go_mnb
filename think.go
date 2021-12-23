package main

import (
	"fmt"
)

func think(ir int, sensor_data_string string) {
	//fmt.Println("\nIR,DATA IN: ", ir, sensor_data_string)

	var brain Brain
	brain = rovers[ir].brain
	var sensor_data [NUM_NEURONS]byte

	sensor_data = convert_sensor_data(sensor_data_string)

	var temp_outps [NUM_NEURONS]byte
	var memb [NUM_NEURONS]int //because memb can go negative
	var outps [NUM_NEURONS]byte
	var fire_knt [NUM_NEURONS]int
	var inps [NUM_NEURONS]byte


	inps = sensor_data
	//fmt.Println("INPS: ", inps)

	sign := brain.sign

	for epoch := 0; epoch < SETTLING_TIME; epoch++ {
		for k := 0; k < NUM_NEURONS; k++ {
			outps[k] = temp_outps[k]
			temp_outps[k] = 0
		}
		//fmt.Println("AT TOP OUTPS: ", outps)

		for nindex := 0; nindex < NUM_NEURONS; nindex++ {
			if outps[nindex] == 0 {
				memb[nindex] = 0
				//not in refactory state
				//do input to membrane
			for ilink:=0;ilink<NUM_NEURONS;ilink++ {
				//iconn would have had a nindex if iconn
				//changed from neuron to neuron
			memb[nindex] += int(inps[ilink] * brain.iconn[ilink])
			}
		}
		} //end of loop on nindex for inputs
		fmt.Println("\nENDOF NINDEX INPUTS")
		fmt.Println("MEMB  : ",memb)

		for nindex := 0; nindex < NUM_NEURONS; nindex++ {
			if outps[nindex] == 0 {
			stuff := 0
			for il := 0; il < NUM_NEURONS; il++ {
			stuff +=  int(outps[il] * brain.nconn[nindex][il]) *sign[il]
			}
			memb[nindex] += stuff
        		}
		} //end of nindex just save membrane and compute firings after

		fmt.Println("\n END OF EPOCH",epoch)
		fmt.Println("MEMB       : ", memb)
		temp_outps = get_output_state(memb)

		//fire_knt is used to choose what sensor to go with
		for k := 0; k < NUM_NEURONS; k++ {
			fire_knt[k] += int(temp_outps[k])
		}
	} //end of settling_time loop (epochs)


	dx := temp_outps[0] + temp_outps[NUM_NEURONS-1]
	dy := temp_outps[1] + temp_outps[NUM_NEURONS-2]
	if dx == 0 {
		rovers[ir].Vel_x = 0
	}
	if dx == 1 {
		rovers[ir].Vel_x = 1
	}
	if dx == 2 {
		rovers[ir].Vel_x = -1
	}
	if dy == 0 {
		rovers[ir].Vel_y = 0
	}
	if dy == 1 {
		rovers[ir].Vel_y = 1
	}
	if dy == 2 {
		rovers[ir].Vel_y = -1
	}

} //end of think




func convert_sensor_data(sensor_data_string string) [NUM_NEURONS]byte {
	//could have called Atoi on this, but meh
	var sig byte
	sig = 0
	var sensor_data [NUM_NEURONS]byte

	for i := 0; i < len(sensor_data_string); i++ {
		if sensor_data_string[i] == '0' {
			sig = 0
		} else {
			sig = 1
		}
		sensor_data[i] = sig
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

func n_n_membrane(nconn [NUM_NEURONS]byte, sign [NUM_NEURONS]byte, outps [NUM_NEURONS]byte) int {

	junk := 0
	for il := 0; il < NUM_NEURONS; il++ {
		if sign[il] == 1 {
			junk += int(outps[il] * nconn[il])
		} else {
			junk -= int(outps[il] * nconn[il])
		}

	}
	return junk
}

func get_output_state(membrane [NUM_NEURONS]int) [NUM_NEURONS]byte {
	var junk [NUM_NEURONS]byte

	for nindex := 0; nindex < NUM_NEURONS; nindex++ {
		junk[nindex] = 0
		if membrane[nindex] < 0 {
			membrane[nindex] = 0
		}
		if membrane[nindex] >= LEAKING_CONSTANT {
			membrane[nindex] -= LEAKING_CONSTANT
		}

		r := getRandomInt(-2, 3)
		if membrane[nindex] >= (THRES + r) {
			junk[nindex] = 1
		}
	}
	return junk
}
