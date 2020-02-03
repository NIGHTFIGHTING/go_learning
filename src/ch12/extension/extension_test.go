package extension

import (
    "testing"
    "fmt"
)
// -------------------------Pet-----------------------
type Pet struct {
}
func (p *Pet) Speak() {
    fmt.Print("...")
}
// SpeakeTo调用Speak函数
func (p *Pet) SpeakTo(host string) {
    p.Speak();
    fmt.Println(" ", host)
}
// -------------------------Dog扩展Pet-----------------------
type Dog struct {
    p *Pet
}
func (d *Dog) Speak() {
    // d.p.Speak();
    fmt.Print("wang")
}
func (d *Dog) SpeakTo(host string) {
    // 如果只更改Dog的Speak(),这里调用不能使用Dog的Speak函数
    // d.p.SpeakTo(host)
    // 需要使用Dog的Speak函数需要为下面的方式
    d.Speak();
    fmt.Println(" ", host)
}
func TestDog(t *testing.T) {
    dog := new(Dog)
    // 如果不修改SpakeTo函数，只修改Speak函数，则打印的为：...  chao，因此需要修改SpeakTo函数
    // 不能使用Dog的Speak函数,使用的是Pet的Speak函数.期望打印wang chao,因此需要修改SpeakTo函数
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
