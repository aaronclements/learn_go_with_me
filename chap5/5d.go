package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	262.6
//
//  Comment:			This is a callback to 5c.  But we do it differently.
//                      We use the "_" because we can't declare a variable we won't use, like "i".
//                      The character is a way of saying "we don't need this."
//
//                      Also, do we really need to use so many lines to declare our array?
//                      Here, I say no.  We can use as many lines as we want, but it can depend.
//
//						( 98 + 233 + 77 + 822 + 83 ) = 1313
//                      len(x) = 5
//						( 1313 / len(x) ) = 262.6
//

func main() {
	x: [5]float64{ 98, 233, 77, 822, 83 }

    var total float64 = 0
    for _, value := range x {
        total += value
    }
    fmt.Println( total / float64(len(x) ) )
}