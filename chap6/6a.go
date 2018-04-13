package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	46.2
//
//						We're creating "xs", which contains 22,33,77,44,55.
//						"total" is set to 0.0, making it a float
//						"v" starts at 0.0 and "total" takes that and adds each value in array "xs"
//						"total" is then divided by the length of the array, giving us an average.

func main() {
	xs := []float64{22,33,77,44,55}

	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println(total / float64(len(xs)))
}
