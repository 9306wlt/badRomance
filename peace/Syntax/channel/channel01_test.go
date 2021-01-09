package channel

import (
    "fmt"
    "testing"
)

func send01(ch chan interface{}){
    for i:=1;i<20;i++{
        ch <- i
        fmt.Printf("sendData = %v \n",i)
    }
    close(ch)
}

func receive(ch chan interface{},exit chan bool){
    for{
        v,ok := <- ch
        if !ok{
            break
        }
        fmt.Printf("receiveData = %v \n",v)
    }
    exit <- true
    close(exit)
}

func Test_channel01(t *testing.T){
    ch := make(chan  interface{},1)
    exit := make(chan bool)
    go send01(ch)

    go receive(ch,exit)

    for{
        _,ok := <-exit
        if !ok{
            break
        }
    }

}
