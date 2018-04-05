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
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}

// 
// Create a variable named i with the value 1.
// Is i <= 10? Yes.
// Print i.
//
// Set i to i + 1 (i now equals 2).
// Is i <= 10? Yes.
// Print i.
//
// Set i to i + 1 (i now equals 3).
//
// ...
//
// Set i to i + 1 (i now equals 11).
// Is i <= 10? No.
// Nothing left to do, so exit.”
//
