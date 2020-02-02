package main

import (
    "fmt"
)

func adder() func(int) int {
    sum := 0
    return func(v int) int {
        sum += v
        return sum
    }
}
func TestAdd() {
    a := adder()
    for i := 0; i < 10; i++ {
        fmt.Printf("0 + ... + %d = %d\n", i, a(i))
    }
}
type iAdder func(int) (int, iAdder)
func adder2(base int) iAdder {
    return func(v int) (int, iAdder) {
        return base + v, adder2(base +v)
    }
}

func TestAdd2() {
    a := adder2(0)
    for i := 0; i < 10; i++ {
        var s int
        s, a = a(i)
        fmt.Printf("0 + ... + %d = %d\n", i, s)
    }
}

func main() {
    TestAdd()
    TestAdd2()
}

//0 + ... + 0 = 0
//0 + ... + 1 = 1
//0 + ... + 2 = 3
//0 + ... + 3 = 6
//0 + ... + 4 = 10
//0 + ... + 5 = 15
//0 + ... + 6 = 21
//0 + ... + 7 = 28
//0 + ... + 8 = 36
//0 + ... + 9 = 45
