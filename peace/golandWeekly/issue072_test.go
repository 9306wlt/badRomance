package golandWeekly

import (
    "fmt"
    "testing"
)


func Test(t *testing.T){
    var(
        a int = 0
        b int64 = 0
        c interface{} = int(0)
        d interface{} = int64(0)

    )

    //fmt.Printf("a的数据类型 %T,b的数据类型 %T,c的数据类型 %T,d的数据类型 %T\n",a,b,c,d)
    //fmt.Println(IntToBytes(a))
    //fmt.Println(Int64ToBytes(b))
    if c == b{
        fmt.Println(true)
    }else{
        fmt.Println(false)
    }

    if c == 0{
        fmt.Println(true)
    }else{
        fmt.Println(false)
    }

    if c == a{
        fmt.Println(true)
    }else{
        fmt.Println(false)
    }

    if d == b{
        fmt.Println(true)
    }else{
        fmt.Println(false)
    }


    //println(d == 0)
}
