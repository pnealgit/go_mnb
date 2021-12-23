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
//	"time"
)


var rovers [NUM_ROVERS]Rover
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

func talk(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	//big loop... maybe
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		} //end of if on message err
		fixed_mt = mt

		fmt.Println("MESSAGE TYPE: ", mt)
		junk := string(message)
		fmt.Println("ARENA DATA: ", junk)
		if strings.Contains(junk, "make_arena") {
			fmt.Println("MAKE ARENA!!")
			jerr := json.Unmarshal(message, &arena)
			if jerr != nil {
				fmt.Println("error on arena unmarshal")
				os.Exit(3)
			} //end of if on jerr
			fmt.Println("ARENA WIDTH: ", arena.Width)
			fmt.Println("ARENA HEIGHT: ", arena.Height)
			fmt.Println("ARENA FOOD: ", arena.Food)
			make_rovers()
			break
		} //end of if on arena
	} //end of infinite loop waiting for message

	//ok now we just spew data to web
	var draw_message []byte
	var draw_positions [NUM_ROVERS][2]int
	var mmm Mess
	for try := 0; try < 100; try++ {
		fmt.Println("TRY: ", try)
		for num_steps := 0; num_steps < NUM_MAX_STEPS; num_steps++ {
				draw_positions = do_update()
				dead_knt := 0
        			for ix:= 0;ix<NUM_ROVERS;ix++ {
                			if rovers[ix].Dead {
                        			dead_knt += 1
                			}
				}
				if dead_knt >= NUM_ROVERS {
					dead_knt = 0
					//fmt.Println("DEAD KNT >= NUM_ROVERS")
					break
				}

				mmm.Msg_type = "positions"
				mmm.Positions = draw_positions
				draw_message, err = json.Marshal(mmm)
				if err != nil {
					fmt.Println("bad angles Marshal")
					os.Exit(7)
				}

				err = c.WriteMessage(fixed_mt, draw_message)
				if err != nil {
					log.Println("BAD DRAW MESSAGE:", err)
					os.Exit(4)
				} //end of if on write err
			//time.Sleep(5 * time.Millisecond)
		} //loop on num_steps
		select_brains()
	} //end of try loop
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
