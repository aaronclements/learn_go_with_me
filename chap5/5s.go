package main
import (
	"fmt"
	"sort"
)

// Aaron Clements did this.
//
//  COMMENT NEEDS TO BE CHECKED......
//
//  Expected output:
//
//  Comment:
//
//

func main() {
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
    }

    sort.Ints(x)
    lowest := x[0]
    fmt.Println(lowest)
}