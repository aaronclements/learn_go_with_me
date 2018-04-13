package main

import "fmt"

// Aaron Clements did this.
//
// Expected output:     true
//                      false
//                      true
//                      true
//                      false
//

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
