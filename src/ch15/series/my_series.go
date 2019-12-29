package series


func Square(n int) int {
    return n*n
}

func GetbonacciSeries(n int) [] int {
    //list []int = []int{1, 1}
    list := []int{1, 1}
    for i := 2; i < n; i++ {
        list = append(list, list[i-2] + list[i-1])
    }
    return list
}
