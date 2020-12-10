package timer

import (
    "fmt"
    "testing"
    "time"
)

func Test_ticker(t *testing.T){

    d := time.Duration(time.Second*2)
    ticker := time.NewTicker(d)

    defer ticker.Stop()

    for{
         a := <- ticker.C
         fmt.Printf("a的类型为%T a的值为%v\n",a,a)
         fmt.Println("timeout...")
    }
}
