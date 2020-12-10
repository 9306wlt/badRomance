package 常见语法题一

import (
    "fmt"
    "testing"
)

type Param map[string]interface{}

type Show struct {
    Param
}

func Test_question01(t *testing.T){
    //这一步会报错，new关键字无法初始化Show结构体中的Param属性。
    s := new(Show)
    s.Param["RMB"] = 10000
    fmt.Println(s)


    //解决方案一
    /*
       s := Show{Param{"RMB":10000}}
       fmt.Println(s)
     */

    //解决方案二
    /*
       s := new(Show)
       s.Param = make(Param，0）
       s.Param["RMB"] = 10000
       fmt.Println(s)
    */
}
