package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	7th	
//			5th
//
//  Comment:		fifth()    will print "5th"
//			seventh()  will print "7th"		
//
//			main()     will print seventh() first
//			           because fifth() has been deferred
//			           until after the function is complete
//
//

func fifth() {
    fmt.Println("5th")
}

func seventh() {
    fmt.Println("7th")
}

func main() {
    defer fifth()
    seventh()
}
