package main

import (
//	"fmt"
	"math"
)

func get_sensor_data(ir int) {
	//the return is a vector of 1s and 0s

	//gotta do it this way to avoid jumping over obstacles
	wall := 0
	var Xpos int
	var Ypos int
	var sensor_angle_index int
	var dist int
	for isensor := 0; isensor < NUM_SENSORS; isensor++ {
		sensor_angle_index = get_sensor_angle_index(isensor, rovers[ir].Angle_index)
		//sensor_angle_index = isensor
		deltax := ANGLES_DX[sensor_angle_index]
		deltay := ANGLES_DY[sensor_angle_index]
		Xpos = rovers[ir].Xpos
		Ypos = rovers[ir].Ypos
		wall = 0
		for step := 0; step < SENSOR_LENGTH; step++ {
			Xpos += deltax
			Ypos += deltay
			fdx := float64(Xpos - rovers[ir].Xpos)
			fdy := float64(Ypos - rovers[ir].Ypos)
			dist = int(math.Hypot(fdx, fdy))
			if dist >= SENSOR_LENGTH {
				dist = SENSOR_LENGTH
				break
			}
			wall = check_wall_position(Xpos, Ypos)
			if wall > 0 {
				break
			}

			//0 means is a sensor not a rover
			wall = check_food_position(Xpos, Ypos,0)
			if wall > 0 {
				break
			}
		} //end of step loop
		rovers[ir].Sensor_data[isensor][0] = Xpos
		rovers[ir].Sensor_data[isensor][1] = Ypos
		rovers[ir].Sensor_data[isensor][2] = wall
		rovers[ir].Sensor_data[isensor][3] = dist
	} //end of isensor loop

	//rovers[ir] = rovers[ir]
}

func make_binary_sensor_data(ir int) string {
	//fmt.Println("in make binary : ",rovers[ir].Sensor_data)

	//knt := 0
	//obstacles
	//hmmm lets just do good/bad on wall/food
	//no distances
	//good := "11"
	//bad  := "00"
	//zip  := "01"

	nothing := "0001"
	no_go := "0011"
	eats := "1111"

	//distances
	far := "1000"
	soso := "1001"
	clos := "1011" //close is a reserved word
	alert := "1011"

	var bsd string
	bsd = ""

	//string contains 3 groups containing a type code and a distance code
	for i := 0; i < NUM_SENSORS; i++ {
		otype := rovers[ir].Sensor_data[i][2]
		if otype == 0 {
			bsd = bsd + nothing
			//	bsd = bsd + zip
		}

		if otype > 0 && otype < 7 {
			//bsd = bsd + bad
			bsd = bsd + no_go
		}
		if otype == 7 {
			//bsd = bsd + good
			bsd = bsd + eats
		}

		dist := rovers[ir].Sensor_data[i][3]
		junkf := float64(dist)
		slf := float64(SENSOR_LENGTH)
		if junkf > .8*slf {
			if otype > 0 {
				bsd = bsd + far
			} else {
				bsd = bsd + nothing
			}
			continue
		}
		//if junkf > .5*slf && junkf <= .8*slf {
		if junkf > .5*slf {
			bsd = bsd + soso
			continue
		}
		//if junkf > .15*slf && junkf <= .5 {
		if junkf > .15*slf {
			bsd = bsd + clos
			continue
		}
		//if junkf <= .15*slf {
		bsd = bsd + alert
		//}
	}

	//fmt.Println("IN BINARY STRIN: ",bsd)
	return bsd
}

func check_wall_position(xp int, yp int) int {

	if yp <= 2 {
		return 1
	}

	if yp >= arena.Height-2 {
		return 2
	}

	if xp <= 2 {
		return 3
	}
	if xp >= arena.Width-2 {
		return 4
	}
	return 0

}

func check_food_position(xp int, yp int,isRover int) int {
	status := 0

	for i := 0; i < len(prey); i++ {
		if prey[i].Dead == 1 {
			continue
		}
		dx := float64(prey[i].Xpos - xp)
		dy := float64(prey[i].Ypos - yp)
		dist := int(math.Hypot((dx), (dy)))
		if dist <= FOOD_RADIUS {
			if isRover == 1{
				prey[i].Dead = 1
				//fmt.Println("PREY ",i," IS DEAD")

			}
			status = 7
			break
		}
	}
	return status
}

func get_sensor_angle_index(isensor int, rover_angle_index int) int {
	ai := 99
	//first sensor
	if isensor == 0 {
		ai = rover_angle_index + 1
	}

	//middle sensor points in direction of movement
	if isensor == 1 {
		ai = rover_angle_index
	}

	if isensor == 2 {
		ai = rover_angle_index - 1
	}

	if ai > NUM_ANGLES-1 {
		ai = ai % NUM_ANGLES
	}

	if ai < 0 {
		ai = NUM_ANGLES - 1
	}

	return ai
}
