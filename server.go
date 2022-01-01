package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"strings"
	//		"time"
)

var rovers [NUM_ROVERS]Rover
var prey []Prey

var arena Arena
var num_steps int
var addr = flag.String("addr", "localhost:8081", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var fixed_mt int

var draw_message []byte
var rover_position [8]int
var prey_positions [][3]int
var mmm Mess
var this_rover int
func talk(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("upgrade:", err)
		return
	}
	defer c.Close()
	//big loop... maybe
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		} //end of if on message err
		fixed_mt = mt

		junk := string(message)
		//fmt.Println("ARENA DATA: ", junk)
		if strings.Contains(junk, "make_arena") {
			fmt.Println("MAKE ARENA!!")
			jerr := json.Unmarshal(message, &arena)
			if jerr != nil {
				fmt.Println("error on arena unmarshal")
				os.Exit(3)
			} //end of if on jerr
			fmt.Println("ARENA WIDTH: ", arena.Width)
			fmt.Println("ARENA HEIGHT: ", arena.Height)
			make_rovers()
			make_prey()
			this_rover = 0
		} //end of if on arena

		if strings.Contains(junk, "ACK") {
			num_dead := 0
			num_dead = num_dead_prey()

                       rovers[this_rover].Time_to_live--
                        if rovers[this_rover].Dead == 1 ||
                                rovers[this_rover].Time_to_live <= 0 ||
				num_dead == len(prey) {
                                this_rover++
                                reset_prey()
				if this_rover >= NUM_ROVERS {
					select_brains()
					make_prey()
					this_rover = 0
				}
                        } 
			if this_rover < NUM_ROVERS {
				rover_position = do_rover_update(this_rover)
	 			prey_positions = do_prey_updates()

				mmm.Msg_type = "positions"
				mmm.Predator_position = rover_position
				mmm.Prey_positions = prey_positions

				draw_message, err = json.Marshal(mmm)
				if err != nil {
					fmt.Println("bad angles Marshal")
					os.Exit(7)
				}
				err = c.WriteMessage(fixed_mt, draw_message)
				if err != nil {
					fmt.Println("BAD DRAW MESSAGE:", err)
					os.Exit(4)
				} //end of if on write err
			} 

                } //end of if on ACK
	} //end of infinite loop waiting for message

	fmt.Println("END OF TRY LOOP")
	os.Exit(0)
} //end of talk

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/talk", talk)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	fmt.Println("listening on 8081")
	log.Fatal(http.ListenAndServe(*addr, nil))
} //end of main
