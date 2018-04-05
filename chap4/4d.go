package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:  1
//                   2
//                   3
//                   4
//                   5
//                   6
//                   7
//                   8
//                   9
//                   10

func main() {
	for i := 1 ; i <= 10 ; i++ {
		fmt.Println(i)
	}
}
