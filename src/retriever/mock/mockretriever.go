package mock

import (
    "fmt"
)

// mock.Retriever现在是RetrieverPoster接口
type Retriever struct {
    Contents string
}

// 系统接口
func (r *Retriever) String() string {
    return fmt.Sprintf(
        "Retriever: {Contents=%s}", r.Contents)
}

//func (r *Retriever)Get(url string) string {
func (r Retriever)Get(url string) string {
    return r.Contents 
}

// 这里需要是r *Retriever,否则原对象的r.Contents数值不会改变
func (r *Retriever)Post(url string,
    form map[string]string) string {
    r.Contents = form["contents"]
    return "ok"
}
