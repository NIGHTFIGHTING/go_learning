package extension

import (
    "testing"
    "fmt"
)
type Pet struct {
}
func (p *Pet) Speak() {
    fmt.Print("...")
}
func (p *Pet) SpeakTo(host string) {
    p.Speak();
    fmt.Println(" ", host)
}
type Dog struct {
    p *Pet
}
func (d *Dog) Speak() {
    // d.p.Speak();
    fmt.Print("wang")
}
func (d *Dog) SpeakTo(host string) {
    // 如果只更改Dog的Speak(),这里调用不能使用Dog的Speak函数
    //d.p.SpeakTo(host)
    d.Speak();
    fmt.Println(" ", host)
}
func TestDog(t *testing.T) {
    dog := new(Dog)
    // ...  chao
    // 不能使用Dog的Speak函数,使用的是Pet的Speak函数.期望打印wang chao
    dog.SpeakTo("chao");
}

// ------------------这个地方是cat
type Cat struct {
    Pet
}

func (c *Cat) Speak() {
    fmt.Print("wang")
}

func TestCat(t *testing.T) {
    cat := new(Cat)
    cat.SpeakTo("chao")
}
