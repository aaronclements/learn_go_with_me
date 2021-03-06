package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	Aaron Clements
//
//  Comment:			Here, we are element mapping strings to strings.
//                      In this case, we map "SRE" to another string.
//                      It then returns the string mapped to the SRE element.

func main() {
    SRE := make(map[string]string)

    SRE["AC"]  = "Aaron Clements"
    SRE["JS"]  = "Josh Stigall"
    SRE["DZ"]  = "Daniel Zhang"
    SRE["EBD"] = "Eamon Bisson-Donahue"
    SRE["MA"]  = "Michael Anzuoni"
    SRE["AW"]  = "Alan Wilson"
    SRE["PT"]  = "Pete Teitelbaum"
    SRE["SL"]  = "Sean Lyons"
    SRE["ZS"]  = "Zack Stayman"

    fmt.Println(SRE["AC"])
}