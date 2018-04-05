package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:  Enter number of feet: 6.2
//                   That is 1.88976 in meters.
//

func main() {
	fmt.Print("Enter number of feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	output  := input * 0.3048

	fmt.Println( "That is", output, "in meters." )
}
