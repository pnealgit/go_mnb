package main
import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"time"
)
	        type Rover struct {
                X       int
                Y       int
                Fitness int
                Z       [5]int
	}
func main() {

	start := time.Now()
// some computation
elapsed := time.Since(start)
fmt.Println(elapsed)

	var rover Rover
rover.X = 500
rover.Y = 300
rover.Fitness = 50
rover.Z = [5]int{1,2,3,4,5}

fmt.Println("WRITIING BEST ")
//fmt.Println("GLOBAL FITNESS: ",GLOBAL_FITNESS," BEST ROVER FITNESS: ",rover[0].Fitness)
GLOBAL_FITNESS := 99

fmt.Println("GLOBAL FITNESS: ",GLOBAL_FITNESS," BEST ROVER FITNESS: ",rover.Fitness)
	/*f,err := os.Create("best.txt")
	if err != nil {
		fmt.Println("bad create")
		os.Exit(20)
	}
	defer f.Close()
*/
	//var best_message []bytes

	best_message, err := json.Marshal(rover)
        if err != nil {
              fmt.Println("bad best rover Marshal")
              os.Exit(7)
        }
	fmt.Println("BEST MESSAGE: ",best_message)


	err = ioutil.WriteFile("best_rover.json", best_message, 0644)
	if err != nil {
		fmt.Println("BAD WRITEFILE")
		fmt.Println(err)
	}

	content, err := ioutil.ReadFile("best_rover.json")
	if err != nil {
		fmt.Println("BAD READ")
		os.Exit(21)
	}
	var best_rover Rover

	err = json.Unmarshal(content, &best_rover)
	if err != nil {
		fmt.Println("BAD UNMARSHAL",err)
		os.Exit(22)
	}
	fmt.Println("BEST ROVER: ",best_rover)
	fmt.Printf("X: %4d Y: %4d",best_rover.X,best_rover.Y)
}
