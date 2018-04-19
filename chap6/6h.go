package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:    5 6
//
//	Comment:	We can place multiple values in a function.
//			Here, we store two int's, 5 and 6.
//			Then we grab the values from f1 into x and y.
//			Then we print the values of x and y.
//			

func f1() (int, int) {
	return 5, 6
}

func main() {
	x, y := f1()
	fmt.Println(x,y)
}
