package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	5
//
//		Functions don't have access to anything unless it's passed exlicitly.
//

func f(x int) {
	fmt.Println(x)
}

func main() {
	x := 5
	f(x)
}