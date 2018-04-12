package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	[1 2 3 4 5 6 7 8 9 10] [0 0 0 0 0]
//
//  Comment:			We are setting x=5 and y=9 for this example.
//                      slice1 is explicitly made as [1 2 3 4 5 6 7 8 9 10]
//                      slice2 has a capacity of y ints, defaulted to 0's
//                      slice2 is [0 0 0 0 0 0 0 0 0]  (9 zeroes)
//                      slice2 has room for x integers (5 integers)
//                      only the first x values of array y will be printed
//                     
//

func main() {
    x := 5
    y := 9
    slice1 := []int{1,2,3,4,5,6,7,8,9,10}
    slice2 := make([]int, x, y)
    fmt.Println(slice1, slice2)
}
