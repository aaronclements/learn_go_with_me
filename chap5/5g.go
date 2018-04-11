package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	[1 2 3] [1 2] [1 2]
//
//  Comment:			slice1 is [1 2 3]
//                      slice2 only has room for two integers
//                      slice2 starts out as [0 0]
//                      values are copied to slice2 from slice1
//                      but slice2 only has room for the first two
//                      slice2 is now [1 2]
//

func main() {
    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2,slice1)
    fmt.Println(slice1, slice2, slice2)
}