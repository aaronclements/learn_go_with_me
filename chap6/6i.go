package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:    6
//
//	Comment:	Here, we create a function called "add".
//			"total" starts at 0. We will add all "g"
//			values we define in main() to "total". 
//
//			main() uses add() to perform "1 + 2 + 3"
//			and returns the value "6".
//
//			add() can be called with multiple integers. 
//			This is known as a variadic parameter.


func add(args ...int) int {
    total := 0
    for _, g := range args {
        total += g
    }
    return total
}

func main() {
    fmt.Println(add(1,2,3))
}

