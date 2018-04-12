package main
import "fmt"

// Aaron Clements did this.
//
//	Task:				Access the 4th element of an array.
//
//  Expected output:	100
//
//  Comment:			in an array of 7 x's
//						the fourth x is set to 100
//						ints are else set to 0 by default
//						this will return 100

func main() {
	var x [7]int
	x[4] = 100
	fmt.Println(x[4])
}
