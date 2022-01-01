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
	    //get things started
    	junk = {}
    	junk["msg_type"] = "ACK";
    	//console.log("ACK MESSAGE: ",junk);
    	senddata(junk);
	
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
		ROVER = response.Predator_position;
		//console.log("ROVER IS: ",ROVER)


    		myGameArea.clear();
	    draw_rover();

		if (response.Prey_positions !== null) { 
			var prey = response.Prey_positions;
		      r = 10
        ctx = myGameArea.context;
        for(var j=0;j<prey.length;j++ ) {
        ctx.beginPath();
        ctx.fillStyle = "red";
        ctx.arc(prey[j][0],prey[j][1], r, 0, 2 * Math.PI);
        ctx.fill();
        ctx.stroke();
	}
		} //end of if
    	junk = {}
    	junk["msg_type"] = "ACK";
    	//console.log("ACK MESSAGE: ",junk);
    	senddata(junk);
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
    myGameArea.clear();
    //draw_rovers();
    //draw_prey();
} //end of updateGameArea

myGameArea = {
    canvas : document.createElement("canvas"),
    start : function() {
        //this.millis = 75;  //game intervale milliseconds
        this.canvas.width = width;
        this.canvas.height = height;
        this.context = this.canvas.getContext("2d");
        document.body.insertBefore(this.canvas, document.body.childNodes[0]);
        pause = false;
        //this.interval = setInterval(updateGameArea,this.millis);
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

