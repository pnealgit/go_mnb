package main

//remember kids, Marshal only converts members of a struct if
//the name is Capitalized.... 55 minutes on that problem...
type Mess struct {
        Msg_type  string
        Predator_position [8]int
        Prey_positions [][3]int
        //Position [8]int
}

type Arena struct {
        Width  int
        Height int
//        Food   [][2]int
        Epochs int
}

type Prey struct {
	Xpos	int
	Ypos	int
	Dead	int
}

type Brain struct {
        seed  int64
        sign  [NUM_NEURONS]int
        iconn [NUM_NEURONS]byte
        nconn [NUM_NEURONS][NUM_NEURONS]byte
}


//as of Dec 19, the only sensor data I have
//to hang on to are the end positions of each sensor
//and that is only for drawing purposes
type Rover struct {
        Luts       [NUM_NEURONS][8]int
        Xpos        int
        Ypos        int
        Fitness     int
        Vel_x           int
        Vel_y           int
        Sensor_data [NUM_SENSORS][4]int
        Dead        int
	Angle_index int
	Time_to_live int
}





const NUM_SENSORS = 3
var SENSOR_LENGTH = 160 

//gotta have 1 neuron per binary digit
//4 digits for distance code
//4 digits for type code
//so for now NUM_NEURONS should equal 24

//NUM_SENSORS * 4 * 4
const NUM_NEURONS = 64
const NUM_ROVERS = 20
const STATE_SIZE = 128
var INPS_SIZE = 3
var SETTLING_TIME = 10
var MUTATION_RATE = .2    //"Use the fucking float, Luke"

//e,ne,n,nw,w,sw,s,se
var ANGLES_DX = [8]int{1, 1,  0, -1, -1, -1, 0, 1}
var ANGLES_DY = [8]int{0,-1, -1, -1,  0,  1, 1, 1}
var NUM_ANGLES = 8
var FOOD_RADIUS = 10
var NUM_MAX_STEPS = 2000
var NUM_TRIES = 500
var GLOBAL_FITNESS = 0
var BOX_WIDTH = 21
var BOX_HEIGHT = 21
var BOX_HALF   = 10
