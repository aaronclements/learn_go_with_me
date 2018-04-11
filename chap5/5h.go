package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	map[trackID:9746]
//
//  Comment:			we make "x" a "map" of strings to ints
//

func main() {
    x := make(map[string]int)
    x["trackID"] = 9746
    fmt.Println(x)
}