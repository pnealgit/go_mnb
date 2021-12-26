function Prey(xy) {
	
    this.x = xy.shift();
    this.y = xy.shift();
    this.r = 15;
    this.color = 'green';

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

function draw_prey() {
//	console.log("DRAWING LEN ",PREY.length)
	color = "green"
	r = 10
        for(var j=0;j<PREY.length;j++ ) {
		//console.log("DRAW Px,Py",PREY[j].x,PREY[j].y)
       		ctx = myGameArea.context;
            //ctx.fillStyle = "green"
	    ctx.fillStyle = '#F9DC5C';

            //ctx.beginPath();
		ctx.fillRect(PREY[j][0],PREY[j][1],20,20)
            //ctx.arc(PREY[j].x,PREY[j].y,r,0,2*Math.PI);
            //ctx.fill();
            //ctx.strokeStyle = '#ff0000';
            //ctx.stroke();
            //ctx.closePath();
       }
}
