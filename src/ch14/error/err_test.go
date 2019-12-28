package err_test

import (
    "testing"
    "errors"
    "fmt"
    "strconv"
)

var LessThanTwoError = errors.New("n shoulde be not less than 2")
var LargerThanHundredError = errors.New("n shoulde be not larger than 100")

func GetFibonacci(n int) ([] int,error) {
    if n < 2 {
        return nil, LessThanTwoError 
    }
    if n > 100 {
    }
    fibList := []int{1, 1}

    for i := 2; i < n ; i++ {
        fibList = append(fibList, fibList[i-2] + fibList[i-1])
    }
    return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
    if v, err := GetFibonacci(1); err != nil {
        if err == LessThanTwoError {
            fmt.Println("It is less")
        }
        t.Error(err)
    } else {
        t.Log(v)
    }
}

func GetFibonacci2(n string) {
    var (
        num int
        err error
        list []int
    )
    // 避免嵌套
    if num, err = strconv.Atoi(n); err != nil {
        fmt.Println("Error", err)
        return
    }
    if list, err = GetFibonacci(num); err != nil {
        fmt.Println("Error", err)
        return
    }
    fmt.Println(list)
}

func TestGetFibonacci2(t *testing.T) {
    GetFibonacci2("1")
}
