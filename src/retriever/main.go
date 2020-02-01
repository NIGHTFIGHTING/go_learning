package main

import (
    "fmt"
    "retriever/mock"
    "retriever/real"
    "time"
)

const url = "http://www.imooc.com"
// 下载接口
type Retriever interface {
    Get(url string) string
}

func download(r Retriever) string {
    return r.Get(url)
}

// 上传接口
type Poster interface {
    Post(utl string, form map[string]string) string
}

func post(poster Poster) {
    poster.Post(url,
        map[string]string {
            "name": "ccmouse",
            "course": "golang",
        })
}

// 可以有多个小接口组合
type RetrieverPoster interface {
    Retriever
    Poster
}
func session(s RetrieverPoster) string {
    s.Post(url, map[string]string {
        "contents": "another faked imooc.com",
    })
    return s.Get(url)
}

func main() {
    var r Retriever
    // mock/mockretriever.go
    // mock.Retriever现在是RetrieverPoster接口
    fmt.Println("-----------测试mock.Retriever接口----------------")
    r = &mock.Retriever{Contents:"this is a fake imooc.com"}
    inspect(r)
    // fmt.Println(download(r))
    //Type assertion
    mockRetriever := r.(*mock.Retriever)
    fmt.Println("Type assertion: "+mockRetriever.Contents)

    fmt.Println("\n------------测试real.RealRetriever接口---------------")

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
    if mockRetriever,ok := r.(*mock.Retriever); ok {
        fmt.Println("Type assertion: "+mockRetriever.Contents)
    } else {
        fmt.Println("not mock a retriever")
    }
    fmt.Println("\n----------测试RetrieverPoster接口-----------------")
    fmt.Println(session(mockRetriever))
}

// 如何判断接口的类型
func inspect(r Retriever) {
    // %T查看类型 %v查看数值
    fmt.Printf(">类型:%T 数值:%v\n", r, r)
    fmt.Printf(">Type switch:")
    switch v := r.(type) {
    case *mock.Retriever:
        fmt.Println("Contents: "+v.Contents)
    case *real.RealRetriever:
        fmt.Println("UserAgent: "+v.UserAgent)
    }
    //mock.Retriever {this is a fake imooc.com}
    //Contents: this is a fake imooc.com
    //*real.RealRetriever &{Mozilla/5.0 1m0s}
    //UserAgent: Mozilla/5.0
}
