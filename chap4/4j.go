package main
import "fmt"

// Aaron Clements did this.
//
// “Write a program that prints out all the numbers between 1 and 100
// that are evenly divisible by 3 (i.e., 3, 6, 9, etc.)"
//

func main() {
    for i := 1; i <= 100; i++ {
        if i % 3 == 0 {
            fmt.Println(i, "is evenly divisible by 3!")
        } else {
            fmt.Println(i, "is not evenly divisible by 3. Sorry.")
        }
    }
    for i := 1; i <= 100; i++ {
        if i % 6 == 0 {
            fmt.Println(i, "is evenly divisible by 6!")
        } else {
            fmt.Println(i, "is not evenly divisible by 6. Sorry.")
        }
    }
    for i := 1; i <= 100; i++ {
        if i % 9 == 0 {
            fmt.Println(i, "is evenly divisible by 9!")
        } else {
            fmt.Println(i, "is not evenly divisible by 9. Sorry.")
        }
    }
}

/*  Expected results:
        1 is not evenly divisible by 3. Sorry.
        2 is not evenly divisible by 3. Sorry.
        3 is evenly divisible by 3!
        4 is not evenly divisible by 3. Sorry.
        5 is not evenly divisible by 3. Sorry.
        6 is evenly divisible by 3!
        7 is not evenly divisible by 3. Sorry.
        8 is not evenly divisible by 3. Sorry.
        9 is evenly divisible by 3!
        10 is not evenly divisible by 3. Sorry.
        11 is not evenly divisible by 3. Sorry.
        12 is evenly divisible by 3!
        13 is not evenly divisible by 3. Sorry.
        14 is not evenly divisible by 3. Sorry.
        15 is evenly divisible by 3!
        16 is not evenly divisible by 3. Sorry.
        17 is not evenly divisible by 3. Sorry.
        18 is evenly divisible by 3!
        19 is not evenly divisible by 3. Sorry.
        20 is not evenly divisible by 3. Sorry.
        21 is evenly divisible by 3!
        22 is not evenly divisible by 3. Sorry.
        23 is not evenly divisible by 3. Sorry.
        24 is evenly divisible by 3!
        25 is not evenly divisible by 3. Sorry.
        26 is not evenly divisible by 3. Sorry.
        27 is evenly divisible by 3!
        28 is not evenly divisible by 3. Sorry.
        29 is not evenly divisible by 3. Sorry.
        30 is evenly divisible by 3!
        31 is not evenly divisible by 3. Sorry.
        32 is not evenly divisible by 3. Sorry.
        33 is evenly divisible by 3!
        34 is not evenly divisible by 3. Sorry.
        35 is not evenly divisible by 3. Sorry.
        36 is evenly divisible by 3!
        37 is not evenly divisible by 3. Sorry.
        38 is not evenly divisible by 3. Sorry.
        39 is evenly divisible by 3!
        40 is not evenly divisible by 3. Sorry.
        41 is not evenly divisible by 3. Sorry.
        42 is evenly divisible by 3!
        43 is not evenly divisible by 3. Sorry.
        44 is not evenly divisible by 3. Sorry.
        45 is evenly divisible by 3!
        46 is not evenly divisible by 3. Sorry.
        47 is not evenly divisible by 3. Sorry.
        48 is evenly divisible by 3!
        49 is not evenly divisible by 3. Sorry.
        50 is not evenly divisible by 3. Sorry.
        51 is evenly divisible by 3!
        52 is not evenly divisible by 3. Sorry.
        53 is not evenly divisible by 3. Sorry.
        54 is evenly divisible by 3!
        55 is not evenly divisible by 3. Sorry.
        56 is not evenly divisible by 3. Sorry.
        57 is evenly divisible by 3!
        58 is not evenly divisible by 3. Sorry.
        59 is not evenly divisible by 3. Sorry.
        60 is evenly divisible by 3!
        61 is not evenly divisible by 3. Sorry.
        62 is not evenly divisible by 3. Sorry.
        63 is evenly divisible by 3!
        64 is not evenly divisible by 3. Sorry.
        65 is not evenly divisible by 3. Sorry.
        66 is evenly divisible by 3!
        67 is not evenly divisible by 3. Sorry.
        68 is not evenly divisible by 3. Sorry.
        69 is evenly divisible by 3!
        70 is not evenly divisible by 3. Sorry.
        71 is not evenly divisible by 3. Sorry.
        72 is evenly divisible by 3!
        73 is not evenly divisible by 3. Sorry.
        74 is not evenly divisible by 3. Sorry.
        75 is evenly divisible by 3!
        76 is not evenly divisible by 3. Sorry.
        77 is not evenly divisible by 3. Sorry.
        78 is evenly divisible by 3!
        79 is not evenly divisible by 3. Sorry.
        80 is not evenly divisible by 3. Sorry.
        81 is evenly divisible by 3!
        82 is not evenly divisible by 3. Sorry.
        83 is not evenly divisible by 3. Sorry.
        84 is evenly divisible by 3!
        85 is not evenly divisible by 3. Sorry.
        86 is not evenly divisible by 3. Sorry.
        87 is evenly divisible by 3!
        88 is not evenly divisible by 3. Sorry.
        89 is not evenly divisible by 3. Sorry.
        90 is evenly divisible by 3!
        91 is not evenly divisible by 3. Sorry.
        92 is not evenly divisible by 3. Sorry.
        93 is evenly divisible by 3!
        94 is not evenly divisible by 3. Sorry.
        95 is not evenly divisible by 3. Sorry.
        96 is evenly divisible by 3!
        97 is not evenly divisible by 3. Sorry.
        98 is not evenly divisible by 3. Sorry.
        99 is evenly divisible by 3!
        100 is not evenly divisible by 3. Sorry.
        1 is not evenly divisible by 6. Sorry.
        2 is not evenly divisible by 6. Sorry.
        3 is not evenly divisible by 6. Sorry.
        4 is not evenly divisible by 6. Sorry.
        5 is not evenly divisible by 6. Sorry.
        6 is evenly divisible by 6!
        7 is not evenly divisible by 6. Sorry.
        8 is not evenly divisible by 6. Sorry.
        9 is not evenly divisible by 6. Sorry.
        10 is not evenly divisible by 6. Sorry.
        11 is not evenly divisible by 6. Sorry.
        12 is evenly divisible by 6!
        13 is not evenly divisible by 6. Sorry.
        14 is not evenly divisible by 6. Sorry.
        15 is not evenly divisible by 6. Sorry.
        16 is not evenly divisible by 6. Sorry.
        17 is not evenly divisible by 6. Sorry.
        18 is evenly divisible by 6!
        19 is not evenly divisible by 6. Sorry.
        20 is not evenly divisible by 6. Sorry.
        21 is not evenly divisible by 6. Sorry.
        22 is not evenly divisible by 6. Sorry.
        23 is not evenly divisible by 6. Sorry.
        24 is evenly divisible by 6!
        25 is not evenly divisible by 6. Sorry.
        26 is not evenly divisible by 6. Sorry.
        27 is not evenly divisible by 6. Sorry.
        28 is not evenly divisible by 6. Sorry.
        29 is not evenly divisible by 6. Sorry.
        30 is evenly divisible by 6!
        31 is not evenly divisible by 6. Sorry.
        32 is not evenly divisible by 6. Sorry.
        33 is not evenly divisible by 6. Sorry.
        34 is not evenly divisible by 6. Sorry.
        35 is not evenly divisible by 6. Sorry.
        36 is evenly divisible by 6!
        37 is not evenly divisible by 6. Sorry.
        38 is not evenly divisible by 6. Sorry.
        39 is not evenly divisible by 6. Sorry.
        40 is not evenly divisible by 6. Sorry.
        41 is not evenly divisible by 6. Sorry.
        42 is evenly divisible by 6!
        43 is not evenly divisible by 6. Sorry.
        44 is not evenly divisible by 6. Sorry.
        45 is not evenly divisible by 6. Sorry.
        46 is not evenly divisible by 6. Sorry.
        47 is not evenly divisible by 6. Sorry.
        48 is evenly divisible by 6!
        49 is not evenly divisible by 6. Sorry.
        50 is not evenly divisible by 6. Sorry.
        51 is not evenly divisible by 6. Sorry.
        52 is not evenly divisible by 6. Sorry.
        53 is not evenly divisible by 6. Sorry.
        54 is evenly divisible by 6!
        55 is not evenly divisible by 6. Sorry.
        56 is not evenly divisible by 6. Sorry.
        57 is not evenly divisible by 6. Sorry.
        58 is not evenly divisible by 6. Sorry.
        59 is not evenly divisible by 6. Sorry.
        60 is evenly divisible by 6!
        61 is not evenly divisible by 6. Sorry.
        62 is not evenly divisible by 6. Sorry.
        63 is not evenly divisible by 6. Sorry.
        64 is not evenly divisible by 6. Sorry.
        65 is not evenly divisible by 6. Sorry.
        66 is evenly divisible by 6!
        67 is not evenly divisible by 6. Sorry.
        68 is not evenly divisible by 6. Sorry.
        69 is not evenly divisible by 6. Sorry.
        70 is not evenly divisible by 6. Sorry.
        71 is not evenly divisible by 6. Sorry.
        72 is evenly divisible by 6!
        73 is not evenly divisible by 6. Sorry.
        74 is not evenly divisible by 6. Sorry.
        75 is not evenly divisible by 6. Sorry.
        76 is not evenly divisible by 6. Sorry.
        77 is not evenly divisible by 6. Sorry.
        78 is evenly divisible by 6!
        79 is not evenly divisible by 6. Sorry.
        80 is not evenly divisible by 6. Sorry.
        81 is not evenly divisible by 6. Sorry.
        82 is not evenly divisible by 6. Sorry.
        83 is not evenly divisible by 6. Sorry.
        84 is evenly divisible by 6!
        85 is not evenly divisible by 6. Sorry.
        86 is not evenly divisible by 6. Sorry.
        87 is not evenly divisible by 6. Sorry.
        88 is not evenly divisible by 6. Sorry.
        89 is not evenly divisible by 6. Sorry.
        90 is evenly divisible by 6!
        91 is not evenly divisible by 6. Sorry.
        92 is not evenly divisible by 6. Sorry.
        93 is not evenly divisible by 6. Sorry.
        94 is not evenly divisible by 6. Sorry.
        95 is not evenly divisible by 6. Sorry.
        96 is evenly divisible by 6!
        97 is not evenly divisible by 6. Sorry.
        98 is not evenly divisible by 6. Sorry.
        99 is not evenly divisible by 6. Sorry.
        100 is not evenly divisible by 6. Sorry.
        1 is not evenly divisible by 9. Sorry.
        2 is not evenly divisible by 9. Sorry.
        3 is not evenly divisible by 9. Sorry.
        4 is not evenly divisible by 9. Sorry.
        5 is not evenly divisible by 9. Sorry.
        6 is not evenly divisible by 9. Sorry.
        7 is not evenly divisible by 9. Sorry.
        8 is not evenly divisible by 9. Sorry.
        9 is evenly divisible by 9!
        10 is not evenly divisible by 9. Sorry.
        11 is not evenly divisible by 9. Sorry.
        12 is not evenly divisible by 9. Sorry.
        13 is not evenly divisible by 9. Sorry.
        14 is not evenly divisible by 9. Sorry.
        15 is not evenly divisible by 9. Sorry.
        16 is not evenly divisible by 9. Sorry.
        17 is not evenly divisible by 9. Sorry.
        18 is evenly divisible by 9!
        19 is not evenly divisible by 9. Sorry.
        20 is not evenly divisible by 9. Sorry.
        21 is not evenly divisible by 9. Sorry.
        22 is not evenly divisible by 9. Sorry.
        23 is not evenly divisible by 9. Sorry.
        24 is not evenly divisible by 9. Sorry.
        25 is not evenly divisible by 9. Sorry.
        26 is not evenly divisible by 9. Sorry.
        27 is evenly divisible by 9!
        28 is not evenly divisible by 9. Sorry.
        29 is not evenly divisible by 9. Sorry.
        30 is not evenly divisible by 9. Sorry.
        31 is not evenly divisible by 9. Sorry.
        32 is not evenly divisible by 9. Sorry.
        33 is not evenly divisible by 9. Sorry.
        34 is not evenly divisible by 9. Sorry.
        35 is not evenly divisible by 9. Sorry.
        36 is evenly divisible by 9!
        37 is not evenly divisible by 9. Sorry.
        38 is not evenly divisible by 9. Sorry.
        39 is not evenly divisible by 9. Sorry.
        40 is not evenly divisible by 9. Sorry.
        41 is not evenly divisible by 9. Sorry.
        42 is not evenly divisible by 9. Sorry.
        43 is not evenly divisible by 9. Sorry.
        44 is not evenly divisible by 9. Sorry.
        45 is evenly divisible by 9!
        46 is not evenly divisible by 9. Sorry.
        47 is not evenly divisible by 9. Sorry.
        48 is not evenly divisible by 9. Sorry.
        49 is not evenly divisible by 9. Sorry.
        50 is not evenly divisible by 9. Sorry.
        51 is not evenly divisible by 9. Sorry.
        52 is not evenly divisible by 9. Sorry.
        53 is not evenly divisible by 9. Sorry.
        54 is evenly divisible by 9!
        55 is not evenly divisible by 9. Sorry.
        56 is not evenly divisible by 9. Sorry.
        57 is not evenly divisible by 9. Sorry.
        58 is not evenly divisible by 9. Sorry.
        59 is not evenly divisible by 9. Sorry.
        60 is not evenly divisible by 9. Sorry.
        61 is not evenly divisible by 9. Sorry.
        62 is not evenly divisible by 9. Sorry.
        63 is evenly divisible by 9!
        64 is not evenly divisible by 9. Sorry.
        65 is not evenly divisible by 9. Sorry.
        66 is not evenly divisible by 9. Sorry.
        67 is not evenly divisible by 9. Sorry.
        68 is not evenly divisible by 9. Sorry.
        69 is not evenly divisible by 9. Sorry.
        70 is not evenly divisible by 9. Sorry.
        71 is not evenly divisible by 9. Sorry.
        72 is evenly divisible by 9!
        73 is not evenly divisible by 9. Sorry.
        74 is not evenly divisible by 9. Sorry.
        75 is not evenly divisible by 9. Sorry.
        76 is not evenly divisible by 9. Sorry.
        77 is not evenly divisible by 9. Sorry.
        78 is not evenly divisible by 9. Sorry.
        79 is not evenly divisible by 9. Sorry.
        80 is not evenly divisible by 9. Sorry.
        81 is evenly divisible by 9!
        82 is not evenly divisible by 9. Sorry.
        83 is not evenly divisible by 9. Sorry.
        84 is not evenly divisible by 9. Sorry.
        85 is not evenly divisible by 9. Sorry.
        86 is not evenly divisible by 9. Sorry.
        87 is not evenly divisible by 9. Sorry.
        88 is not evenly divisible by 9. Sorry.
        89 is not evenly divisible by 9. Sorry.
        90 is evenly divisible by 9!
        91 is not evenly divisible by 9. Sorry.
        92 is not evenly divisible by 9. Sorry.
        93 is not evenly divisible by 9. Sorry.
        94 is not evenly divisible by 9. Sorry.
        95 is not evenly divisible by 9. Sorry.
        96 is not evenly divisible by 9. Sorry.
        97 is not evenly divisible by 9. Sorry.
        98 is not evenly divisible by 9. Sorry.
        99 is evenly divisible by 9!
        100 is not evenly divisible by 9. Sorry.
*/