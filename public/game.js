var ws
var pause = false

function WebsocketStart() {
	console.log("WEBSOCKET START")
      	//setup();
    ws = new WebSocket("ws://localhost:8081/talk")
    ws.onopen = function(evt) {
	junk = {}
	junk['Width'] = width;
	junk['Height'] = height;
	junk['Epochs'] = 100;
	junk['msg_type'] = "make_arena";	
        senddata(junk);
      	myGameArea.start(); 
    }
    ws.onclose = function(evt) {
      console.log('WEBSOCKET CLOSE');
      myGameArea.stop();
      //ws = null;
    }

    ws.onmessage = function(e) {
      	n = e.data.indexOf("position");
      	if (n != -1 ) {
	 	var response = JSON.parse(e.data)
		//console.log("RESPONSE: ",response)
	      	if (ROVERS.length <= 0) {
			make_new_rovers(response.Predator_positions);
	      	} else {
			update_rovers(response.Predator_positions)
		}

	      	if (PREY.length <= 0) {
			make_new_prey(response.Prey_positions);
	      	} else {
			update_prey(response.Prey_positions)
		}

		//console.log("PREY AFTER LOAD: ",PREY)

		
      } //end of if found 'positions'
    } //endo of onmessage


    ws.onerror = function(evt) {
        console.log('onerror ',evt.data);
    }

} //end of WebsocketStart

senddata = function(data) {
    if (pause) {
      return;
    }
    if (!ws) {
        console.log('cannot send data -- no ws');
        return false;
    }
    stuff = JSON.stringify(data);
    ws.send(stuff);
} //end of function senddata

    
function updateGameArea() {
    if (pause) {
       return
    }
       var mydata = {};
       //mydata['num_episodes'] =  num_episodes;
       //senddata(mydata);

    myGameArea.clear();
    draw_rovers();
    draw_prey();
} //end of updateGameArea

myGameArea = {
    canvas : document.createElement("canvas"),
    start : function() {
        this.millis = 75;  //game intervale milliseconds
        this.canvas.width = width;
        this.canvas.height = height;
        this.context = this.canvas.getContext("2d");
        document.body.insertBefore(this.canvas, document.body.childNodes[0]);
        pause = false;
        this.interval = setInterval(updateGameArea,this.millis);
    },  
    stop : function() {
        pause = true; 
        console.log("STOP !!! ");
        clearInterval(this.interval);
        //ws.close();
    },  
    clear : function() {
        this.context.clearRect(0, 0, this.canvas.width, this.canvas.height);
        this.context.fillStyle = "rgba(255,255,255,255)";
        this.context.fillRect(0,0,this.canvas.width,this.canvas.height);
    } 
}    //end of gamearea

