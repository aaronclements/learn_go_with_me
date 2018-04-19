package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	0
//
//  Comment:	The use of the pointer here allows x to be modified
//		outside of its own function, unlike 6q.
//
// 		Pointers reference a location in memory where a value is stored
//		rather than the value itself. By using a pointer (*int), the
//		zero function is able to modify the original variable.	
//

func zero(xPtr *int) {
    *xPtr = 0
}

func main() {
    x := 5
    zero(&x)
    fmt.Println(x)
}
