package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:    10
//			11
//			12
//
//	Comment:	Here, x starts as 9.
//			"increment" takes x and adds 1 each time it's called.
//			Here, we print it twice.
//			So we are taking x=9, adding 1 and printing x, now 10.
//			Then we take x=10, add 1 and print 11.	
//
//

func main() {
    x := 9			// x = 9 
    increment := func() int {
        x++			// now x = 10 and is returned as "increment"
        return x		
    }
    fmt.Println(increment())	// x is returned as 10 when it is called
    fmt.Println(increment())	// since "increment" ran again, now we get 11.
    fmt.Println(increment())	// and now we get 12...
}
