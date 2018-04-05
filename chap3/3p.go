package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:  Enter a number: 
//                   5
//                   30
//

func main() {
	fmt.Println( "Enter a number: " )
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2 * 3
	fmt.Println(output)
}
