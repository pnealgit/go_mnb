//flipped from squares to circles
function Prey(xy) {
	
    this.x = xy.shift();
    this.y = xy.shift();
this.dead = xy.shift();
    this.r = 10;
    this.color = 'red';

} //end of prey function  

function make_new_prey(positions) {
     for (var pos=0;pos < positions.length;pos++) {
           PREY[pos] = new Prey(positions[pos]);
	     //console.log("NEW POS,PREY POS",pos,PREY[pos])

     }
}
function update_prey(positions) {
    for (var pos=0;pos < positions.length;pos++) {
        PREY[pos] = positions[pos];
	     //console.log("OLD POS,PREY POS",pos,PREY[pos])
    }
} //end of function

function draw_prey(prey) {
	r = 10
       	ctx = myGameArea.context;
        for(var j=0;j<prey.length;j++ ) {
		//if (prey[j][2] == 0) {
	ctx.beginPath();
	ctx.fillStyle = "red";

		//flipped
	ctx.arc(prey[j][0],prey[j][1], r, 0, 2 * Math.PI);
	ctx.fill();
	ctx.stroke();
	//}
       }
}
