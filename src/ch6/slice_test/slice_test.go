package slice_test

import "testing"

func TestSlicaInit(t *testing.T) {
    var s0 []int
    t.Log(len(s0), cap(s0))
    s0 = append(s0, 1)
    t.Log(len(s0), cap(s0))

    s1 := []int{1,2,3,4}
    t.Log(len(s1), cap(s1))

    s2 := make([]int, 3, 5)
    t.Log(len(s2), cap(s2))
    // panic: runtime error: index out of range [recovered]
    // t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
    t.Log(s2[0], s2[1], s2[2])
    s2 = append(s2, 1)
    t.Log(s2[0], s2[1], s2[2], s2[3])
    t.Log(len(s2), cap(s2))
}

// 切片如何实现可变长
// 存储空间cap2倍增长
//slice_test.go:29: 1 1
//slice_test.go:29: 2 2
//slice_test.go:29: 3 4
//slice_test.go:29: 4 4
//slice_test.go:29: 5 8
//slice_test.go:29: 6 8
//slice_test.go:29: 7 8
//slice_test.go:29: 8 8
//slice_test.go:29: 9 16
//slice_test.go:29: 10 16
func TestSlicGrowing(t *testing.T) {
    s := []int{}
    for i := 0; i < 10; i++ {
        // 为什么重新复制为s
        // 因为连续存储空间会变化,存储重建复制
        s = append(s, i)
        t.Log(len(s), cap(s))
    }
}

func TestSliceShareMemory(t *testing.T) {
    year := []string{"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep",
        "Oct","Nov","Dec"}
    Q2 := year[3:6]
    t.Log(Q2,len(Q2), cap(Q2))

    summer := year[5:8]
    t.Log(summer, len(summer), cap(summer))
    summer[0] = "Unkonw"
    t.Log(Q2)
    t.Log(year)
}

func TestSliceCompareing(t *testing.T) {
    a := []int{1,2,3,4}
    b := []int{1,2,3,4} 
    // ./slice_test.go:62:10: invalid operation: a == b (slice can only be compared to nil)
    // if a == b {
    //    t.Log("equal")
    // }
}

// 数组和切片的不同
//1.数组容量不可变
//2.数组可以比较
