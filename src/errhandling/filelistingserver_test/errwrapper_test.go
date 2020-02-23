package filelistingserver_test 
//package main


import (
    "fmt"
    "errors"
    "os"
    "testing"
    "net/http"
    "net/http/httptest"
    "strings"
    "io/ioutil"
    "errhandling/filelistingserver_test/global"
)
//// 实现type testingUserError interface接口
//type testingUserError string
//
//// type error interface {
////     Error() string
//// }
//// 实现error的接口
//func (e testingUserError) Error() string {
//    return e.Message()
//}
//
//func (e testingUserError) Message() string {
//    return string(e)
//}

func errPanic(writer http.ResponseWriter,
    request *http.Request) error {
    panic(123)
}

func errUserError(writer http.ResponseWriter,
    request *http.Request) error {
    return global.TestingUserError("user error")
}

func errNotFound(writer http.ResponseWriter,
    request *http.Request) error {
    return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter,
    request *http.Request) error {
    return os.ErrPermission
}

func errUnkonwn(writer http.ResponseWriter,
    request *http.Request) error {
    return errors.New("unknown error")
}

func noError(writer http.ResponseWriter,
    request *http.Request) error {
    fmt.Fprintln(writer, "no error")
    return nil 
}
var tests = []struct {
    h global.AppHandler
    code int
    message string
}{
    {errPanic, 500, "Internal Server Error"},
    {errUserError, 400, "user error"},
    {errNotFound, 404, "Not Found"},
    {errNotPermission, 403, "Forbidden"},
    {errUnkonwn, 500, "Internal Server Error"},
    {noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {

    for _, tt := range tests {
        f := global.ErrWrapper(tt.h)
        response := httptest.NewRecorder()
        request := httptest.NewRequest(
            http.MethodGet,
            "http://www.imooc.com"/*target*/,nil/*body*/,
        )
        f(response, request)
        verifyResponse(response.Result(), tt.code, tt.message, t)
    }
}


// 真正启动一个服务器测试
func TestErrWrapperInServer(t *testing.T) {
    for _, tt := range tests {
        f := global.ErrWrapper(tt.h)
        server := httptest.NewServer(
            http.HandlerFunc(f))
        resp, _ := http.Get(server.URL)
        verifyResponse(resp, tt.code, tt.message, t)
    }
}

func verifyResponse(resp *http.Response,
    expectedCode int, expectedMsg string,
    t *testing.T) {
    b, _ := ioutil.ReadAll(resp.Body)
    body := strings.Trim(string(b), "\n")
    if resp.StatusCode != expectedCode ||
        body != expectedMsg {
        t.Errorf("expect (%d, %s); " +
            "got (%d, %s)",
            expectedCode, expectedMsg,
            resp.StatusCode, body)
    }
}
