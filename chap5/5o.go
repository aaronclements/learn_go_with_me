package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	Josh Stigall
//
//  Comment:			This example shows how we can store multiple values in a map.
//                      So we print the first and last name of the requested initials.
//
//                      We now have a map of strings to maps of strings to strings.
//                      The outer map is used as a lookup for the SRE table.
//                      The inner maps store the information about the SRE's.

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

    if name, ok := SRE["JS"]; ok {
        fmt.Println( name["firstname"], name["lastname"] )
    }
}