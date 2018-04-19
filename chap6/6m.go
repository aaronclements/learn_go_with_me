package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	10
//			16
//			22
//			28	
//
//  Comment:		We create xGen, wherein g starts at 10.
//			xGen returns another function, which says
//			xyz starts at g and has 6 added to it each
//			time it is called.
//
//			
//
//
func xGen() func() uint {
    g := uint(10)
    return func() (xyz uint) {
        xyz = g 
        g += 6 
        return
    }
}
func main() {
    nextOne := xGen()
    fmt.Println(nextOne()) // 10
    fmt.Println(nextOne()) // 16
    fmt.Println(nextOne()) // 22
    fmt.Println(nextOne()) // 28
}
