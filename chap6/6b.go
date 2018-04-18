package main

import "fmt"

// Aaron Clements did this.
//
//  Expected output:	46.2
//
//		Same idea.  "averagetime" takes a slice of float64s and
// 		returns one float64 after dividing by the length of the array.
//		We get the return value from "averagetime" and use it in "main".
//

func averagetime(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {
	xs := []float64{22, 33, 77, 44, 55}
	fmt.Println(averagetime(xs))
}
