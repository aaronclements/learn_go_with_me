package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:    6
//
//	Comment:	Let's encase this function inside a function.
//

func main() {
    add := func(x, y, z int) int {
        return x + y + z
    }
    fmt.Println(add(1,2,3))
}

