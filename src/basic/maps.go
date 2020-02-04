package main

import (
    "fmt"
)

func main() {
    m := map[string]string {
    "name": "ccmouse",
        "course": "golang",
        "site": "imooc",
        "quality": "notbad",
    }

    m2 := make(map[string]int) // m2 == empty
    var m3 map[string]int // m3 == nil
    fmt.Println(m, m2, m3)

    for k, v := range m {
        fmt.Println(k, v)
    }
    for _, v := range m {
        fmt.Println(v)
    }
    for k := range m {
        fmt.Println(k)
    }

    // 打印的k v是无序的，所以map是hash map
    // 如果需要对map k进行排序，将k取出来放到slice中，然后在遍历map得到排序
    fmt.Println("Getting values")
    courseName, ok := m["course"]
    fmt.Println(courseName, ok)
    causeName, ok := m["cause"]
    // 打印了一个空行
    fmt.Println(causeName, ok)

    if causeName, ok := m["cause"]; ok {
        fmt.Println(causeName)
    } else {
        fmt.Println("key does not exist")
    }

    fmt.Println("Deleting values")
    name, ok := m["name"]
    fmt.Println(name, ok)

    delete(m, "name")
    name, ok = m["name"]
    fmt.Println(name, ok)
}
