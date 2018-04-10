package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:  Three
//

func main() {
    var i int
    i = 3
    if i == 0 {
        fmt.Println("Zero")
    } else if i == 1 {
        fmt.Println("One")
    } else if i == 2 {
        fmt.Println("Two")
    } else if i == 3 {
        fmt.Println("Three")
    } else if i == 4 {
        fmt.Println("Four")
    } else if i == 5 {
        fmt.Println("Five")
    }
  }