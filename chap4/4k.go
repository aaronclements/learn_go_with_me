package main
import "fmt"

// Aaron Clements did this.
//
// ““Write a program that prints the numbers from 1 to 100, but...
//      ...for multiples of 3, print “Fizz”.
//      ...for multiples of 5, print “Buzz.”
//      ...for numbers that are multiples of both, print “FizzBuzz.”
//

func main() {
    for i := 1; i <= 100; i++ {
        if i % 3 == 0 {
            fmt.Println(i, "is a Fizz!")
        }
    }
    for i := 1; i <= 100; i++ {
        if i % 5 == 0  {
            fmt.Println(i, "is a Buzz!")
        }
    }
    for i := 1; i <= 100; i++ {
        if i % 5 == 0 && i % 3 == 0 {
            fmt.Println(i, "is a FIZZBUZZ!")
        }
    }
}

/*  Expected output:

3 is a Fizz!
6 is a Fizz!
9 is a Fizz!
12 is a Fizz!
15 is a Fizz!
18 is a Fizz!
21 is a Fizz!
24 is a Fizz!
27 is a Fizz!
30 is a Fizz!
33 is a Fizz!
36 is a Fizz!
39 is a Fizz!
42 is a Fizz!
45 is a Fizz!
48 is a Fizz!
51 is a Fizz!
54 is a Fizz!
57 is a Fizz!
60 is a Fizz!
63 is a Fizz!
66 is a Fizz!
69 is a Fizz!
72 is a Fizz!
75 is a Fizz!
78 is a Fizz!
81 is a Fizz!
84 is a Fizz!
87 is a Fizz!
90 is a Fizz!
93 is a Fizz!
96 is a Fizz!
99 is a Fizz!
5 is a Buzz!
10 is a Buzz!
15 is a Buzz!
20 is a Buzz!
25 is a Buzz!
30 is a Buzz!
35 is a Buzz!
40 is a Buzz!
45 is a Buzz!
50 is a Buzz!
55 is a Buzz!
60 is a Buzz!
65 is a Buzz!
70 is a Buzz!
75 is a Buzz!
80 is a Buzz!
85 is a Buzz!
90 is a Buzz!
95 is a Buzz!
100 is a Buzz!
15 is a FIZZBUZZ!
30 is a FIZZBUZZ!
45 is a FIZZBUZZ!
60 is a FIZZBUZZ!
75 is a FIZZBUZZ!
90 is a FIZZBUZZ!   */