package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:   [ERROR] ./3n.go:11:4: cannot assign to x
//

func main() {
	const x string = "Hello, world."
	x = 5
	fmt.Println( x )
}
