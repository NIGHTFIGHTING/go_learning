package main

import (
    "fmt"
    "errhandling/defer/fib"
    "os"
    "bufio"
)

func tryDefer() {
    defer fmt.Println(1)
    defer fmt.Println(2)
    fmt.Println(3)
    //return
    panic("error occurred")
    fmt.Println(4)
}

func tryDefer1() {
    for i := 0; i < 100; i++ {
        // 参数在defer语句时计算
        defer fmt.Println(i)
        if i == 30 {
            panic("printed too many")
        }
    }
}

func writeFile(filename string) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    // 资源管理
    defer file.Close()

    writer := bufio.NewWriter(file)
    // 资源管理
    defer writer.Flush()

    f := fib.Fibonacci()
    for i := 0; i < 20; i++ {
        fmt.Fprintln(writer, f())
    }
}

func main() {
    //tryDefer()
    //writeFile("fib.txt")
    tryDefer1()
}
