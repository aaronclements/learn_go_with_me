package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	Josh Stigall
//
//  Comment:			This example shows how we can store multiple values.
//

package main
import "fmt"

func main() {
    elements := map[string]map[string]string{
        "AC": map[string]string{
            "firstname":"Aaron",
            "lastname":"Clements",
        },
        "JS": map[string]string{
            "firstname":"Josh",
            "lastname":"Stigall",
        },
    }

    if i, ok := elements["JS"]; ok {
        fmt.Println(i["firstname"], i["lastname"])
    }
}