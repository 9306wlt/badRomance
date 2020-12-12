package 常见语法题二

import (
    "fmt"
    "testing"
)

func defer_call(){
    defer func() {
        fmt.Println("打印前")
    }()
    defer func() {
        fmt.Println("打印中")
    }()
    defer func() {
        fmt.Println("打印后")
    }()

    panic("触发异常")
}

func Test_qeustion01(t *testing.T){
    defer_call()
}
