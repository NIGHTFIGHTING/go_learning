package main 


import (
    "fmt"
)

func swap(a, b *int) {
    *a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
    return b, a
}

func main() {
    a := 3
    b := 4
    swap(&a, &b)
    fmt.Println("a=", a, " b=", b)

    a, b = swap2(a, b)
    fmt.Println("a=", a, " b=", b)
}
