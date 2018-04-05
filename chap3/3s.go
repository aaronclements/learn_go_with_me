package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:  Enter degrees in Farenheit: 85
//                   That is 29.444444444444446 degrees in Celcius.
//

func main() {
	fmt.Print("Enter degrees in Farenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output  := input - 32.0
	output1 := 5.0 / 9.0
	output2 := output * output1

	fmt.Println( "That is", output2, "degrees in Celcius." )
}
