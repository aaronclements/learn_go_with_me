package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:  6
//                   60 
//

func main() {
	x := 5
	x += 1
	fmt.Println( x )
	x *= 10
	fmt.Println( x )
}
