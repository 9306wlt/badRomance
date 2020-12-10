package 常见语法题一

import (
    "fmt"
    "testing"
)

type People4 struct{
    Name string
}

func(p *People4)String() string{
    return fmt.Sprintf("print:%v",p)
    /*
    在使用 fmt 包中的打印方法时，如果类型实现了这个接口，会直接调用。而题目中打印 p 的时候会直接调用 p 实现的 String() 方法，然后就产生了循环调用。
     */
}

func Test_qustion04(t *testing.T){
    p := &People4{}
    p.String()
}
