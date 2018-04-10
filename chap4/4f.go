package main
import "fmt"

// Aaron Clements did this.
//
// Expected output:  1 is odd!
//                   2 is even!
//                   3 is odd!
//                   4 is even!
//                   5 is odd!
//                   6 is even!
//                   7 is odd!
//                   8 is even!
//                   9 is odd!
//                   10 is even!
//

func main() {
    for i := 1; i <= 10; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "is even!")
        } else {
            fmt.Println(i, "is odd!")
        }
    }
}

/*  
    Let’s walk through this program:
        1.  Create a variable i of type int and give it the value 1.
        2.  Is i less than or equal to 10? Yes: jump to the if block.
        3.  Is the remainder of i ÷ 2 equal to 0? No: jump to the else block.
        4.  Print i followed by odd.
        5.  Increment i (the statement after the condition).
        6.  Is i less than or equal to 10? Yes: jump to the if block.
        7.  Is the remainder of i ÷ 2 equal to 0? Yes: jump to the if block.
        8.  Print i followed by even, and so on until i is equal to 11.
        9.  …
*/