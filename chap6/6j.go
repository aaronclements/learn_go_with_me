package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:    6
//
//	Comment:	This passes a slice of ints by following
//			the slice with an elipses...
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
    xs := []int{1,2,3}
    fmt.Println(add(xs...))
}

