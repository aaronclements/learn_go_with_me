package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	262.6
//
//  Comment:			This is a callback to 5b.  But we do it differently.
//                      We will use the length of the array, in case x or i change.
//                      By replacing the "5" with a len(x), we are dividing by the
//                      number of items in the whole array.  It is this way that we
//                      can more confidently say the number returned is a true average.
//
//                      But len(x) is of int type. So we have to encase with float64().
//
//						( 98 + 233 + 77 + 822 + 83 ) = 1313
//                      len(x) = 5
//						( 1313 / len(x) ) = 262.6
//

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 233
	x[2] = 77
    x[3] = 822
    x[4] = 83

    var total float64 = 0
    for i := 0 ; i < len(x) ; i++ {
        total += x[i]
    }
    fmt.Println( total / float64(len(x) ) )
}