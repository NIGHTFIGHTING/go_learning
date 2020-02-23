package main

import (
    "net/http"
    //"os"
    "errhandling/filelistingserver_test/filelisting"
    "errhandling/filelistingserver_test/global"
    //"log"
)

func main() {
    //http.HandleFunc("/list/errhandling/filelistingserver_test/web.go",
    http.HandleFunc("/",
       global.ErrWrapper(filelisting.HandleFileList))
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        panic(err)
    }
}
