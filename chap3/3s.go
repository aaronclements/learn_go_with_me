package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:  6
//                   60 
//

func main() {
	fmt.Print("Enter degrees Farenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output  := input - 32.0
	output1 := 5.0 / 9.0
	output2 := output * output1

	fmt.Println( output2, "degrees Celcius." )
}
