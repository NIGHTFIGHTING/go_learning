package main

import (
    "fmt"
    "retriever/mock"
    "retriever/real"
    "time"
)

type Retriever interface {
    Get(url string) string
}

func download(r Retriever) string {
    return r.Get("http://www.imooc.com")
}

func main() {
    var r Retriever
    // mock/mockretriever.go
    r = mock.Retriever{Contents:"this is a fake imooc.com"}
    inspect(r)
    // fmt.Println(download(r))
    //Type assertion
    mockRetriever := r.(mock.Retriever)
    fmt.Println("Type assertion: "+mockRetriever.Contents)

    fmt.Println("---------------------------")

    // real/retriever.go
    r = &real.RealRetriever{
        UserAgent: "Mozilla/5.0",
        TimeOut: time.Minute,
    }
    inspect(r)
    // fmt.Println(download(r))
    //Type assertion
    realRetriever := r.(*real.RealRetriever)
    fmt.Println("Type assertion: "+realRetriever.UserAgent)
    //  判断接口的类型
    if mockRetriever,ok := r.(mock.Retriever); ok {
        fmt.Println("Type assertion: "+mockRetriever.Contents)
    } else {
        fmt.Println("not mock a retriever")
    }
}

// 如何判断接口的类型
func inspect(r Retriever) {
    // %T查看类型 %v查看数值
    fmt.Printf("类型:%T 数值:%v\n", r, r)
    fmt.Println("Type switch")
    switch v := r.(type) {
    case mock.Retriever:
        fmt.Println("Contents: "+v.Contents)
    case *real.RealRetriever:
        fmt.Println("UserAgent: "+v.UserAgent)
    }
    //mock.Retriever {this is a fake imooc.com}
    //Contents: this is a fake imooc.com
    //*real.RealRetriever &{Mozilla/5.0 1m0s}
    //UserAgent: Mozilla/5.0
}
