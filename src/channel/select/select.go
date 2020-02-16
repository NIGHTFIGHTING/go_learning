package main

import (
    "fmt"
    "time"
    "math/rand"
)


func generator() chan int {
    out := make(chan int)
    go func() {
        i := 0
        for {
            time.Sleep(
                time.Duration(rand.Intn(1500)) *
                    time.Millisecond)
            out <- i
            i++
        }
    }()
    return out
}
func worker(id int, c chan int) {
    for n := range c {
        //time.Sleep(5 * time.Second)
        time.Sleep(1 * time.Second)
        fmt.Printf("Worker %d received %d\n",
            id, n)
    }
}
// <- chan
// chan <-
func createWorker(id int) chan <- int {
    c := make(chan int)
    go worker(id, c)
    return c
}
func main() {
    //var c1, c2 chan int // c1 and c2 = nil
    var c1, c2 = generator(), generator()
    worker := createWorker(0)

    var values []int
    tm := time.After(10 * time.Second)
    tick := time.Tick(time.Second)
    for {
        var activeWorker chan<- int // activeWorker is nil，下面activeWorker <- activeValue会阻塞
        var activeValue int
        if len(values) > 0 {
            activeWorker = worker
            activeValue = values[0]
        }
        select {
        case n := <-c1:
            values = append(values, n)
        case n := <-c2:
            values = append(values, n)
        case activeWorker <- activeValue:
            values = values[1:] 
        case <- time.After(800 * time.Millisecond):
            // 没有tick情况下表示，800ms内generator没有生成数据
            // 有tick表示，两次select处理时差是否超过800ms
            fmt.Println("timeout")
        case <- tick:
            fmt.Println("queue len = ", len(values))
        case <- tm:
            fmt.Println("bye")
            return
        }
    }
}
