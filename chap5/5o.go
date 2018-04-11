package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	Josh Stigall
//
//  Comment:			This example shows how we can store multiple values in a map.
//                      So we print the first and last name of the requested initials.

func main() {
    SRE := map[string]map[string]string{
        "AC": map[string]string{
            "firstname":"Aaron",
            "lastname":"Clements",
        },
        "JS": map[string]string{
            "firstname":"Josh",
            "lastname":"Stigall",
        },
    }

    if i, ok := SRE["JS"]; ok {
        fmt.Println( i["firstname"], i["lastname"] )
    }
}