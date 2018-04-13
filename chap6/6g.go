package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	1
//
//		main will print f1
//		f1 is an int that returns the value in f2 plus 2
//		f2 is an int that returns r
// 		r is set to 1
//		2 + 1 = 3
//


func main() {
	fmt.Println(f1())
}

func f1() int {
	return f2() + 2
}

func f2() (r int) {
	r = 1
	return
}