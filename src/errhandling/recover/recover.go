package main

import (
    //"errors"
    "fmt"
)

func tryRecover() {
    defer func() {
        r := recover()
        // type assertion
        if err, ok := r.(error); ok {
            fmt.Println("Error occurred:", err)
        } else {
            //panic(r)
            panic(fmt.Sprintf(
                "I don't know what to do: %v", r))
            //panic: 123 [recovered]
                //panic: I don't know what to do: 123
        }
    }()
    //panic(errors.New("this is an error"))
    //b := 0
    //a := 5/b
    //fmt.Println(a)
    //Error occurred: runtime error: integer divide by zero

    panic(123)
}

func main() {
    tryRecover()
}
