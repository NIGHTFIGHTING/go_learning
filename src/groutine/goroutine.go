package main


import (
    "fmt"
)

func main() {
    for i := 0; i < 10; i++ {
        go func(i int) {
            for {
                fmt.Printf("Hello from " +
                    "goroutine %d\n", i)
            }
        }(i)
    }
}
