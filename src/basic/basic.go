package main

import (
    "fmt"
    "math"
    "math/cmplx"
)

func consts() {
    // go语言的常量不要求大写
    const (
        filename = "abc.txt"
        a, b = 3, 4
    )
    var c int
    c = int(math.Sqrt(a*a + b*b))
    fmt.Println(filename, c)
}

func enums() {
    const (
        cpp = iota
        _
        python
        golang
        javascript
    )
    fmt.Println(cpp, javascript, python, golang)

    const (
        b = 1 << (10 * iota)
        kb
        mb
        gb
        tb
        pb
    )
    fmt.Println(b, kb, mb, gb, tb, pb)
}

func euler() {
    //c := 3 + 4i
    //fmt.Println(cmplx.Abs(c))
    fmt.Printf(
        "%.3f",
        cmplx.Exp(1i * math.Pi) + 1)
    fmt.Println(
        cmplx.Exp(1i * math.Pi) + 1)
    fmt.Println(
        cmplx.Pow(math.E, 1i * math.Pi) + 1)
}

func triangle() {
    var a, b int = 3, 4
    var c int
    c = int(math.Sqrt(float64(a*a + b*b)))
    fmt.Println(c)
}

func main() {
    consts()
    enums()
    euler()
    triangle()
}
