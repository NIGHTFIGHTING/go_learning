// 06数据类型

package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
    var a int32 = 1
    var b int64
    // 不支持隐式类型转换
    // ./type_test.go:9:7: cannot use a (type int) as type int64 in assignment
    // b = a
    b = int64(a)

    // 不支持别名到原类型类型转换
    var c MyInt
    // c = b
    c = MyInt(b)
    t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
    a := 1
    // 支持指针类型
    aPtr := &a
    // 不支持指针运算
    // aPrt = aPtr + 1
    t.Log(a, aPtr)
    t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
    // 字符串是值类型,默认初始化是空字符串
    var s string
    t.Log("*"+s+"*")
    t.Log(len(s))

    if s == "" {
        t.Log("s is size = 0")
    }
}
