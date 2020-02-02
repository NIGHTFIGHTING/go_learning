package main

import (
    "net/http"
    "os"
    "errhandling/filelistingserver/filelisting"
    "log"
)

type appHandler func(writer http.ResponseWriter,
    request *http.Request) error

// 函数式编程
func errWrapper(
    handler appHandler) func(
        http.ResponseWriter,*http.Request) {
    return func(writer http.ResponseWriter,
        request *http.Request) {
        // 防止访问这种地址出错,http://localhost:8888/li
        // defer + panic + recover
        defer func() {
            if r := recover(); r != nil {
                log.Printf("Panic: %v", r)
                http.Error(writer,
                    http.StatusText(http.StatusInternalServerError),
                    http.StatusInternalServerError)
            }
        }()
        err := handler(writer, request)
        // 错误统一处理
        if err != nil {
            //import "github.com/gpmgo/gopm/modules/log"
            //log.Warn("Error handling request: %s",
              //  err.Error())
            log.Printf("Error handling request: %s",
                err.Error())
            // Type Assertion
            if userErr, ok := err.(userError); ok {
                http.Error(writer,
                    userErr.Message(),
                    http.StatusBadRequest)
                return
            }
            code := http.StatusOK
            switch {
            case os.IsNotExist(err):
                code = http.StatusNotFound
            case os.IsPermission(err):
                code = http.StatusForbidden
            default:
                code = http.StatusInternalServerError
            }
            // 错误都写在writer里
            http.Error(writer,
                    http.StatusText(code),
                    code)
        }
    }
}

// type error interface {
//     Error() string
// }
type userError interface {
    error
    Message() string
}

func main() {
    //http.HandleFunc("/list/errhandling/filelistingserver/web.go",
    http.HandleFunc("/",
       errWrapper(filelisting.HandleFileList))
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        panic(err)
    }
}
