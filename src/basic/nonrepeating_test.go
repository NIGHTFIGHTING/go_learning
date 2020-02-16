package main

import (
    "testing"
)
func lengthOfNonRepeatingSubStr(s string) int {
    //lastOccurred := make(map[byte]int)
    lastOccurred := make(map[rune]int)
    start := 0
    maxLength := 0

    //for i, ch := range []byte(s) {
    for i, ch := range []rune(s) {
        if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
            start = lastOccurred[ch] + 1
        }
        if i-start+1 > maxLength {
            maxLength = i - start + 1
        }
        lastOccurred[ch] = i
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
