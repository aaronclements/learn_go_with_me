package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	10, 10
//
//  Comment:			we make "x" a "map" of ints to ints
//                      we can define "x" references
//                      maps don't require starting with 0

func main() {
    x := make(map[int]int)
    x[1] = 10
    x[2] = 10
    fmt.Println( x[1] , x[2] )
}