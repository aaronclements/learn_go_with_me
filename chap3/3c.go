package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:     first
//                      second
//

func main() {
	var x string 
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
}
