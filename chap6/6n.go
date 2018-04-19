package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	24
//
//  Comment:		"factorials" takes x and returns its factorial
//			"main" says that x starts at 4
//			"factorials" will iterate through x*(x-1)
//			until x reaches 0, at which point it's done.
//
//			(x) * (x-1) * (x-1) * (x-1)
//			(4) * (4-1) * (3-1) * (2-1) = 24
//			(4) * ( 3 ) * ( 2 ) * ( 1 ) = 24	
//
//			1! = 1
//			2! = 2 * 1 = 2
//			3! = 3 * 2 * 1 = 6
//			4! = 4 * 3 * 2 * 1 = 24
//

func factorials(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorials(x-1)
}
func main() {
	var x uint
	x = 4 
	fmt.Println(factorials(x))
}
