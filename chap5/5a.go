package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	[0 0 0 0 100]
//

func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x)
}