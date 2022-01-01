
function Rover(xy) {
	//start in the middle
    this.r = 10;
    this.sensor_data = xy;

/*
    this.draw = function() {
	x = this.sensor_data.shift()
	y = this.sensor_data.shift()
        ctx = myGameArea.context;
        ctx.beginPath();
        ctx.arc(x,y, this.r, 0, 2 * Math.PI);
        ctx.fillStyle = "green";
        ctx.fill();
        ctx.beginPath();
        ctx.strokeStyle = '#000000';
        ctx.stroke();
        ctx.closePath();

	      for(i =0 ;i<NUM_SENSORS;i++) {
		xp = this.sensor_data.shift()
		yp = this.sensor_data.shift()
                ctx.beginPath()
                ctx.strokeStyle = '#000000';
                ctx.moveTo(x,y);
                ctx.lineTo(xp,yp);
                ctx.stroke();
                ctx.closePath();
        }       

    } //end of rover draw
    */
}
//end of Rover function

function draw_rover() {
        //x = ROVER.sensor_data.shift()
        //y = ROVER.sensor_data.shift()
        x = ROVER.shift()
        y = ROVER.shift()
        ctx = myGameArea.context;
        ctx.beginPath();
        ctx.arc(x,y, 10, 0, 2 * Math.PI);
        ctx.fillStyle = "green";
        ctx.fill();
        ctx.beginPath();
        ctx.strokeStyle = '#000000';
        ctx.stroke();
        ctx.closePath();

	for(i =0 ;i<NUM_SENSORS;i++) {
                xp = ROVER.shift()
                yp = ROVER.shift()
                ctx.beginPath()
                ctx.strokeStyle = '#000000';
                ctx.moveTo(x,y);
                ctx.lineTo(xp,yp);
                ctx.stroke();
                ctx.closePath();
        }
} //end of function 

