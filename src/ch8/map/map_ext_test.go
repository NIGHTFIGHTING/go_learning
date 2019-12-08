package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T) {
    m1 := map[int]func(op int) int {}
    m1[1] = func(op int) int { return op }
    m1[2] = func(op int) int { return op*op }
    m1[3] = func(op int) int { return op*op*op }
    t.Log(m1[1](2), m1[2](2), m1[3](2))
}

func TestMapForSet(t *testing.T) {
    mySet := map[int] bool {}
    mySet[1] = true
    // n := 3
    n := 3
    // 判断3是否在set中
    if mySet[n] {
    // if _, ok := mySet[n]; ok {
        t.Logf("%d is existing", n)
    } else {
        t.Logf("%d is not existing", n)
    }
    // set中添加3
    mySet[3] = true
    // 获得set的元素个数
    t.Log("mySet count is", len(mySet))
    // set中删除1
    delete(mySet, 1)
    n = 1
    // 判断1是否在set中
    if mySet[n] {
        t.Logf("%d is existing", n)
    } else {
        t.Logf("%d is not existing", n)
    }
}
