package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	5
//
//		Sema as 6e.  We're making the function main call f() which returns 5.
//

var x int = 5
func f() {
	fmt.Println(x)
}

func main() {
	f()
}