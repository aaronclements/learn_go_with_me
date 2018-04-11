package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	262.6
//
//  Comment:			There is an an array of 5 x's
//						"var total" gets set to 0 and then adds each x.
//						It then prints that sum divided by 5 - also the number of x's.
//						So in this case, we are explicitly calling for an average.
//
//						( 98 + 233 + 77 + 822 + 83 ) = 1313
//						( 1313 / 5 ) = 262.6
//

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 233
	x[2] = 77
    x[3] = 822
    x[4] = 83

    var total float64 = 0
    for i := 0; i < 5; i++ {
        total += x[i]
    }
    fmt.Println(total / 5)
}