package main
import (
	"fmt"
)

func main() {
	fmt.Println("DUH")
	var x [4][2]int

	for i:=0;i<4;i++ {
		for j:=0;j<2;j++{
			x[i][j] = i
		}
	}
	fmt.Println("BEFORE: ",x)

	x[2] = x[1]

	x[2][1] = 99

	x[1][1] =  x[3][1]
	fmt.Println("AFTER:  ",x)

	var z [6]int
	var y [6]int

	for i:=0;i<6;i++ {
		z[i] = i+6
	}
	y = z
	fmt.Println("BEFORE Z: ",z)
	fmt.Println("BEFORE Y: ",y)

	y[3] = 99
	z[2] = y[3]

	y[3] = 77
	fmt.Println("AFTER Z: ",z)
	fmt.Println("AFTER Y: ",y)
}
