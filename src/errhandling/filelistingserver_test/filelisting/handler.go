package filelisting

import (
    "net/http"
    "os"
    "io/ioutil"
    "strings"
    //"errors"
)

const prefix = "/list/"

// 实现type userError interface接口
type userError string

// type error interface {
//     Error() string
// }
// 实现error的接口
func (e userError) Error() string {
    return e.Message()
}

func (e userError) Message() string {
    return string(e)
}

func HandleFileList(writer http.ResponseWriter,
    request *http.Request) error {
    // 请求的url不是/list/开头
    // 防止访问这种地址出错,http://localhost:8888/li
    if strings.Index(
        request.URL.Path, prefix) != 0 {
        //return errors.New("path must start " +
          //  "with " + prefix)
        return userError("path must start " +
            "with " + prefix)
    }
    path := request.URL.Path[len(prefix):] // /list/fib.txt
    file, err := os.Open(path)
    if err != nil {
        // 错误太分散
        //panic(err)
        //http.Error(writer,
          //  err.Error(),
            //http.StatusInternalServerError)
        return err
    }
    defer file.Close()
    all, err := ioutil.ReadAll(file)
    if err != nil {
        // 错误太分散
        //panic(err)
        return err
    }
    writer.Write(all)
    return nil
}
