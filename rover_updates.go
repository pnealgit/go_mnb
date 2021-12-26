package main
import (
//	"fmt"
)

func do_rover_updates() [NUM_ROVERS][8]int {

	//var err error
	var binary_sensor_data string
	//var max_index int
	var team_positions [NUM_ROVERS][8]int //2 for pos 6 for sensor pos

	for ir := 0; ir < NUM_ROVERS; ir++ {
		if rovers[ir].Dead 		{
			continue
		}
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
                        rovers[ir].Fitness += 1 //go straight
                }
		rovers[ir].Angle_index = new_angle_index

		rovers[ir].Xpos += ANGLES_DX[new_angle_index]
		rovers[ir].Ypos += ANGLES_DY[new_angle_index]

		//for dumping back to javascript
		team_positions[ir][0] = rovers[ir].Xpos
		team_positions[ir][1] = rovers[ir].Ypos
		//fmt.Println("SDs: ",rovers[ir].Sensor_data)
		knt := 2
		for ss:=0;ss<NUM_SENSORS;ss++ {
			team_positions[ir][knt] = rovers[ir].Sensor_data[ss][0]
			knt++
			team_positions[ir][knt] = rovers[ir].Sensor_data[ss][1]
			knt++
		}
	} //end of loop on rovers
	return team_positions
} //end of do_update

