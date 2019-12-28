package empty_interface

import (
    "testing"
    "fmt"
)

func Dosomthing(p interface{}) {
    if false {
    if i,ok := p.(int); ok {
        fmt.Println("Interger", i)
        return
    }
    if s,ok := p.(string); ok {
        fmt.Println("string", s)
        return
    }
    fmt.Println("Unknow Type")
    }
    switch v := p.(type) {
    case int:
        fmt.Println("Interger", v)
    case string:
        fmt.Println("String", v)
    default:
        fmt.Println("Unknow Type")
    }
}

func TestEmptyInterfaceAssertion(t *testing.T) {
    Dosomthing(10)
    Dosomthing("10")
}
