package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:  1 is odd!
//                   2 is even!
//                   3 is odd!
//                   4 is even!
//                   5 is odd!
//                   6 is even!
//                   7 is odd!
//                   8 is even!
//                   9 is odd!
//                   10 is even!
//

func main() {
    for i := 1; i <= 10; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "is even!")
        } else {
            fmt.Println(i, "is odd!")
        }
    }
}
