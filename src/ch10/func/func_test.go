package func_test

import (
    "testing"
    "math/rand"
    "time"
    "fmt"
)

func returnMultiValues() (int,int) {
    return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
    a,_ := returnMultiValues()
    t.Log(a)
}

// 函数式编程func programing
// 统一统计时间函数
// inner是函数类型
// 返回值是函数类型
func timeSpent(inner func(op int) int) func(op int) int {
    return func(n int) int {
        start := time.Now()
        ret := inner(n)
        fmt.Println("time spent:", time.Since(start).Seconds())
        return ret
    }
}

func slowFun(op int) int {
    time.Sleep(time.Second * 1)
    return op
}

func TestTimeSpent(t *testing.T) {
    tsSF := timeSpent(slowFun)
    t.Log(tsSF(10))
}


// 可变长参数
func Sum(ops ...int) int {
    ret := 0
    for _, op := range ops {
        ret += op
    }
    return ret
}

func TestVarParam(t *testing.T) {
    t.Log(Sum(1,2,3,4))
    t.Log(Sum(1,2,3,4,5))
}


// defer
// 安全释放资源，释放锁
func Clear() {
    fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
    defer Clear()
    fmt.Println("start")
    // panic 后面的代码不可达
    panic("err")
    fmt.Println("End")
}
