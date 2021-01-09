package channel

import (
    "fmt"
    "testing"
    "time"
)

func pump04_(ch chan int){
    for i:=0; ;i++{
        ch <- i
    }
}

func sumk(ch chan int){
    for{
        fmt.Println(<-ch)
    }
}

func Test_gc04_improve(t *testing.T){

    ch1 := make(chan int)
    go pump04_(ch1)
    go sumk(ch1)
    time.Sleep(1e9)

}

