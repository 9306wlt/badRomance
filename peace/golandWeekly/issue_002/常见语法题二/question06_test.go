package 常见语法题二

import (
    "fmt"
    "testing"
)

func calc(index string, a ,b  int)int{
    ret := a + b
    fmt.Println(index,a,b,ret)
    return ret
}

func Test_question06(t *testing.T){
    a := 1
    b := 2

    defer calc("1",a,calc("10",a,b))
    a = 0

    defer calc("2",a,calc("20",a,b))
    b = 1
}
