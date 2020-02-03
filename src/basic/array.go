package main

import (
    "fmt"
)

// 测试数组是值类型
func printArray(arr [5]int) {
    arr[0] = 100
    for i, v := range arr {
        fmt.Printf("index: %d value: %d\n", i, v)
    }
}
// 使用指针可以修改数组
func printArray1(arr *[5]int) {
    arr[0] = 100
    for i, v := range arr {
        fmt.Printf("index: %d value: %d\n", i, v)
    }
}

func main() {
    var arr1 [5]int
    //arr2 := [3]int{1, 3, 5}
    // 让编译器数数组的长度
    arr3 := [...]int{2, 4, 6, 8, 10}
    //var grid [4][5]int

    fmt.Println("printArray(arr1)")
    printArray(arr1)
    fmt.Println("printArray(arr3)")
    printArray(arr3)
    fmt.Println("arr1 and arr3")
    fmt.Println(arr1, arr3)

    fmt.Println("printArray1(arr1)")
    printArray1(&arr1)
    fmt.Println("printArray1(arr3)")
    printArray1(&arr3)
    fmt.Println("arr1 and arr3")
    fmt.Println(arr1, arr3)

    //fmt.Println(arr1, arr2, arr3)
    //fmt.Println(grid)

    if false {
    // 遍历arr3
    for i := 0; i < len(arr3); i++ {
        fmt.Println(arr3[i])
    }

    // 使用range遍历arr3
    for i, v := range arr3 {
        fmt.Println(i, v)
    }
    // 不需要下标，使用下划线表示
    for _, v := range arr3 {
        fmt.Println(v)
    }
    }
}
