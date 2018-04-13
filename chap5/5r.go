package main
import "fmt"

// Aaron Clements did this.
//
//  COMMENT NEEDS TO BE CHECKED......
//
//  Expected output:	[c d e]
//
//  Comment:		Go up to the 5th string "e"
//                      Go back 2 strings to "c".
//                      Print strings including and between "c" and "e".
//                      ????
//

func main() {
    x := [6]string{"a","b","c","d","e","f"}
    fmt.Println(x[2:5])
}