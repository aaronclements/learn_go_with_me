package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:  Big
//                   Small
//

func main() {
    i := 10
    if i >= 10 {
        fmt.Println("Big")
    } else {
        fmt.Println("Small")
    }
    y := 5
    if y >= 10 {
        fmt.Println("Big")
    } else {
        fmt.Println("Small")
    }
}