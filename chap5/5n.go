package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	Nope!
//
//  Comment:			This example goes back to 5m.
//                      Here, we are still element mapping strings to strings.
//                      In this case, we map "SRE" to another string.
//                      It returns the string mapped to the SRE element.
//                      IN THIS CASE, KS does not exist, so it has no value.
//                      "ok" is whether the search was successful for the value requested.
//
//                      We can also see this example represented in a function.
//                      So we get a value from the map and process it according to the rules.
//                      So instead of just getting " false" for output, I'm having it say "Nope!"
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

    if fullname, ok := SRE["KS"]; ok {
        fmt.Println(fullname, ok) 
    } else{
        fmt.Println("Nope!")
    }
}