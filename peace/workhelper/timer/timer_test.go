package timer

import (
    "fmt"
    "testing"
    "time"
)

func Test_timer(t *testing.T){

    d := time.Duration(time.Second*3)

    timer := time.NewTimer(d)
    defer timer.Stop()

    for{
        b := <-timer.C
        fmt.Printf("b的类型为%T b的值为%v\n",b,b)
        fmt.Println("timeout...")
        timer.Reset(time.Second*2)
    }
}
