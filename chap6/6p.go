package main
import "fmt"

// Aaron Clements did this.
//
//  Expected output:	PANIC NOW!	
//
//  Comment:	setting panic value to "PANIC NOW!"
//		setting str to recover() value, which won't happen
//		attempting to print value of str
//
//
//	func main() {
//		panic("PANIC")
//		str := recover() // this will never happen
//		fmt.Println(str)
//	}
//
//	This will fail because the panic above
//	stops the execution of the program.
//	So we will defer the function that fails.
//

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
        panic("PANIC NOW!")
}
