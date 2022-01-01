package main
import (
//	"fmt"
)

func do_rover_update(ir int) [8]int {

	//var err error
	var binary_sensor_data string
	var team_positions [8]int //2 for pos 6 for sensor pos
	zorro := 0

	Xpos := rovers[ir].Xpos
	Ypos := rovers[ir].Ypos

	zorro = check_food_position(Xpos, Ypos,1)
        if zorro == 7 {
                rovers[ir].Fitness += 1
		//fmt.Println("ROVER : ",ir," FIT: ",rovers[ir].Fitness)
        }

	zorro = 0
	zorro = check_wall_position(Xpos,Ypos)
	if zorro > 0 {
		rovers[ir].Dead = 1
	}

	if rovers[ir].Dead == 0 {	
	get_sensor_data(ir)
	binary_sensor_data = make_binary_sensor_data(ir)
	max_index := 0
	max_index = think(ir, binary_sensor_data)
	        
               //start east go counterclockwise
                //sensors go from left to right 0-2
		var new_angle_index int
                new_angle_index = rovers[ir].Angle_index
                if max_index == 0 {
                        new_angle_index = new_angle_index + 1
                }
                if new_angle_index > NUM_ANGLES-1 {
                        new_angle_index = 0
                }
                if max_index == 1 {
                        //do nothing. Just a straight ahead
                }

                if max_index == 2 {
                        if new_angle_index > 0 {
                                new_angle_index = new_angle_index - 1
                        } else {
                                new_angle_index = NUM_ANGLES - 1
                        }
                }

		//why is this done.. reward for going straight
                if new_angle_index == rovers[ir].Angle_index {
                        //rovers[ir].Fitness += 1 //go straight
                }
		rovers[ir].Angle_index = new_angle_index

		rovers[ir].Xpos += ANGLES_DX[new_angle_index]
		rovers[ir].Ypos += ANGLES_DY[new_angle_index]

		//for dumping back to javascript
		team_positions[0] = rovers[ir].Xpos
		team_positions[1] = rovers[ir].Ypos
		//fmt.Println("SDs: ",rovers[ir].Sensor_data)
		knt := 2
		for ss:=0;ss<NUM_SENSORS;ss++ {
			team_positions[knt] = rovers[ir].Sensor_data[ss][0]
			knt++
			team_positions[knt] = rovers[ir].Sensor_data[ss][1]
			knt++
		}
	} //end of if on dead
	//} //end of loop on rovers
	return team_positions
} //end of do_update

