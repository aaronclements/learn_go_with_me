package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	
//
//  Comment:			Here, we are element mapping strings to strings.
//                      In this case, we map "SRE" to another string.
//                      It then returns the string mapped to the SRE element.
//                      IN THIS CASE, EBD does exist. It has an assigned value.
//                      This means fullname will return Eamon, and ok will return true.
//

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

    fullname, ok := SRE["EBD"]
    fmt.Println(fullname, ok)
}