package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:   Hello, world. 
//                    Hello again!
//

var x string = "Hello, world."
var y string = "Hello again!"

func main() {
	fmt.Println( x )
        fmt.Println( y )
}

func f() {
        fmt.Println( x )
	fmt.Println( x )
}

