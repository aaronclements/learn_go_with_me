package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	307.5
//
//  Comment:			This is a callback to 5d.  I hesitate to even point it out.
//                      We can remove something from the array, if we want, *provided*
//                      we adjust length of the array. This allows us to avoid deleting
//                      a line of our code if we didn't want to.  Trailing comma is required.
//
//						( 98 + 233 + 77 + 822 ) = 1230
//                      len(x) = 4
//						( 1313 / len(x) ) = 307.5
//

func main() {
    x := [4]float64{
        98,
        233,
        77,
        822,
        //83,
    }

    var total float64 = 0
    for _, value := range x {
        total += value
    }
    fmt.Println( total / float64(len(x) ) )
}