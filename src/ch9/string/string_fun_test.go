package string_test

import "testing"
import "strings"
import "strconv"

func TestStringFn(t* testing.T) {
    s := "A,B,C"
    parts := strings.Split(s, ",")
    for index, part := range parts {
        t.Log(index, part)
    }
    // A-B-C
    t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
    s := strconv.Itoa(10)
    // str=10
    t.Log("str="+s)
    if i,err := strconv.Atoi("10"); err == nil {
        // 20
        t.Log(i+10)
    }
}
