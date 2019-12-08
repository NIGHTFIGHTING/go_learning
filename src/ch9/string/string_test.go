package string_test

import "testing"

func TestString(t *testing.T) {
    var s string
    t.Log(s) //初始化为默认零值"",空字符串
    s = "hello"
    t.Log(len(s)) // 5
    //s[1] = '3' //string是不可变的byte slice
    //s = "\xE4\xB8\xA5" //可以存储任何二进制数据
    s = "\xE4\xBA\xBB\xFF"
    t.Log(s) // 乱码
    t.Log(len(s)) // 4
    s = "中"
    t.Log(len(s)) //是byte数, 3

    c := []rune(s)
    t.Log("rune c:", c)
    t.Log(len(c)) // 1
    //  t.Log("rune size:", unsafe.Sizeof(c[0]))
    t.Logf("中 unicode %x", c[0])
    t.Logf("中 UTF8 %x", s)
}

func TestTravelString(t *testing.T) {
    s := "中华人民共和国"
    // 按照rune类型打印
    for _,c := range s {
        t.Logf("%[1]c, %[1]d", c)
    }
}
