package main

import (
    "testing"
)

var lastOccurred = make([]int, 0xffff)
func lengthOfNonRepeatingSubStr(s string) int {
    //lastOccurred := make(map[byte]int)
    //lastOccurred := make(map[rune]int)
    // 测试性能优化
    // stores las occurred pos + 1.
    // 0 means not seen
    for i := range lastOccurred {
        lastOccurred[i] = 0
    }
    start := 0
    maxLength := 0

    //for i, ch := range []byte(s) {
    for i, ch := range []rune(s) {
        if lastI := lastOccurred[ch]; lastI > start {
            start = lastI
        }
        if i-start+1 > maxLength {
            maxLength = i - start + 1
        }
        lastOccurred[ch] = i + 1
    }
    return maxLength
}
func TestSubstr(t *testing.T) {
    tests := []struct {
        s string
        ans int
    } {
        // Normal cases
        {"abcabcbb", 3},
        {"pwwkew", 3},
        // Edge cases
        {"", 0},
        {"b", 1},
        {"bbbbbbbb", 1},
        {"abcabcabcd", 4},
        // Chinense cases
        {"这里是慕课网", 6},
        {"一二三二一", 3},
        {"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
    }

    for _, tt := range tests {
        actual := lengthOfNonRepeatingSubStr(tt.s)
        if actual != tt.ans {
            t.Errorf("got %d for input %s; " +
                "expected %d",
                actual, tt.s, tt.ans)
        }
    }
}

func BenchmarkSubstr(b *testing.B) {
    s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
    for i := 0; i < 13; i++ {
        s = s + s
    }
    b.Logf("len(s) = %d", len(s))
    ans := 8
    // 防止生成s计算在时间里
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        actual := lengthOfNonRepeatingSubStr(s)
        if actual != ans {
            b.Errorf("got %d for input %s; " +
                "expected %d",
                actual, s, ans)
        }
    }
}
