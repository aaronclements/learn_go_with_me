package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:  1
//                   2
//                   3
//                   4
//

func main() {
	i := 1
	for i <= 4 {
		fmt.Println(i)
		i = i + 1
	}
}
