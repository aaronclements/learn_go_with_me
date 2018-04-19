package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	0
//			5	
//
//  Comment:	zero() has an int x
//		in the function, x is valued at 0
//		when we use Println, we see 0 
//
//		main() says x is 5
//		zero will not modify the value of x outside of main()
//		when we call zero with our new x and print it
//		we get 5
//

func zero(x int) {
    x = 0
    fmt.Println(x)
}

func main() {
    x := 5
    zero(x)
    fmt.Println(x)
}
