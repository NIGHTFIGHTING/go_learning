package main

import (
    "fmt"
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

func main() {
    fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
    fmt.Println(lengthOfNonRepeatingSubStr("b"))
    fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
}

/*
func main() {
    lastOccurred := make(map[byte]int)

    str := "abccab"
    start := 0
    maxLength := 0
    length := 0

    for i := 0; i < len(str); i++ {
        //if index, ok := lastOccurred[str[i]]; !ok {
        index, ok := lastOccurred[str[i]]
        if !ok {
            length++
            if (length > maxLength) {
                maxLength = length
            }
        } else {
            start = index + 1
        }
        lastOccurred[str[i]] = i
    }
    fmt.Println("matLength: ", maxLength)
}
*/
