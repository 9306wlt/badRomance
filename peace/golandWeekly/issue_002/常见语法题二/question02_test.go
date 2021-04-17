package 常见语法题二

import (
    "fmt"
    "testing"
)

type student struct{
    Name string
    Age int
}

func Test_question02(t *testing.T){
    m := make(map[string]*student)
    stus := []student{
        {Name:"zhou",Age:24},{Name:"li",Age:23},{Name:"wang",Age:22},
    }
    for _,stu :=range stus{
        fmt.Println(stu,&stu)
        m[stu.Name] = &stu
    }
    fmt.Println(m)

    /*
    make(map[string]*student)中的value是一个student指针，在 for _,stu :=range stus 循环的时候，虽然m[stu.Name]key在变，但是m[stu.Name] = &stu
    每次是把*student指向&stu，这个指针指向的东西随着循环而改变，当循环到最后一个的时候，指针指向的是最后的一个地址，也就是map中所有的key的value都指向最后一个变量的值。
     */
}