package 常见语法题一

import (
    "fmt"
    "testing"
)

type Student struct {
    name string
}

func Test_question09(t *testing.T){
/*
    m := map[string]Student{"people":{"zjl"}}
    //下面的编译会报错 因为map是指针传递，无法修改其中的值。
    m["people"].Name = "wyz"

 */
/*
    //结构体作为map的元素时，不能直接赋值给结构体的某个字段，也就是map中的struct中的字段不能够直接寻址。
    //map作为一个封装好的数据结构，由于它底层可能会由于数据扩张而进行迁移，所以拒绝直接寻址,避免产生野指针；
    //map中的key不存在时，赋值语句其实会进行新的k-v值得插入，所以拒绝直接寻址结构体内的字段，以防结构体不存在的时候可能造成的错误；
    //这可能和map的并发不安全型相关
    x = y 这种赋值的方式，你必须知道x的地址，然后才能把值y赋给x.
    但go中的map的value本身是不可寻址的，因为map的扩容的时候，肯跟要做key/val对迁移
    value本身地址是会改变的

 */
    m := map[string]*Student{"people":&Student{"zjl"}}
    m["people"].name = "wyz"
    v,exists := m["people"]
    if exists {
        fmt.Println("people 的名字是：",*v)
    }else{
        fmt.Println("people 的名字不存在")
    }

}

