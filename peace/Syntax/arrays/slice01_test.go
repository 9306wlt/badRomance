package arrays

import (
    "fmt"
    "testing"
)

func sum(a []int)int{
    s := 0
    for i := 0; i < len(a); i++ {
        s += a[i]
    }
    return s
}

func sumSelf(a []int)[]int{
    s := 1
    for i := 0; i < len(a); i++ {
        a[i] += s
    }
    //fmt.Println("a := ",a)
    return a
}

func Test_slice01(t *testing.T){
    var arr = [5]int{0,1,2,3,4}
    //fmt.Println(sum(arr[:]))
    fmt.Println("初始的arr = ",arr)
    fmt.Println(sumSelf(arr[:]))
    fmt.Println("被引用后的arr = ",arr)
}
