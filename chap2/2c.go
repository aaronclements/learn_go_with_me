package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:     12
//                      101
//                      Hello, World
//

func main() {
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")
}
