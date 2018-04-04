package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:     first 
//                      first second
//

func main() {
	var x string 
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
}
