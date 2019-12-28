package panic_recover

import (
    "testing"
    "errors"
    "fmt"
    //"os"
)

func TestPanicVxExit(t *testing.T) {
    defer func() {
        // 不检测什么错误,记录一下,忽略掉
        //可能是系统某些核心资源消耗完,强制恢复掉,系统依然不能正常工作
        //可能导致系统健康检查程序不能检查楚问题（health check）
        //有些health check是检查进程是否还在，因为进程还在，不能提供服务，僵尸进程
        if err := recover(); err != nil {
            fmt.Println("recovered from", err)
        }
    }()
    fmt.Println("Start")
    //os.Exit(-1)
    panic(errors.New("Something wrong!"))
    fmt.Println("End")
}
