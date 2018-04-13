package main

import "fmt"

// Aaron Clements did this.
//
//  Expected output:	1
//
//		main will print f1
//		f1 is an int that returns the value in f2
//		f2 is an int that returns a value of 1
//

func main() {
	fmt.Println(f1())
}

func f1() int {
	return f2()
}

func f2() int {
	return 1
}