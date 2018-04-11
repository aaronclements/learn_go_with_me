package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	[1 2 3] [1 2 3 4 5]
//
//  Comment:			'slice1' is "1 2 3"
//                      'slize2' is slice2 with "4 5" appended
//                      We then print both on the same line.
//
//

func main() {
    slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)
}