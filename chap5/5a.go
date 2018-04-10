package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	[0 0 10 0 100 0 0]
//
//  Comment:			in an array of 7 x's
//							the fifth x is 100
//							the third x is 10
//							everything else is 0

func main() {
    var x [7]int
	x[4] = 100
	x[2] = 10
	fmt.Println(x)
}