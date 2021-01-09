package channel

import (
	"fmt"
	"testing"
)


func writeData(intChan chan int){
	for i:=1;i<20;i++{
		intChan <- i
		fmt.Printf("writeData = %v \n",i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool){
	for{
		v,ok := <- intChan
		if !ok{
			break
		}
		fmt.Printf("readData = %v \n",v)
	}
	exitChan <- true
	close(exitChan)
}

func Test_goroutine_channel(t *testing.T){

	//创建两个管道，缓冲信道
	//缓冲信到和无缓冲信道，是看你创建信道的时候有没有设置容量
	intChan := make(chan int,50)
	exitChan := make(chan bool,1)

	go writeData(intChan)
	go readData(intChan,exitChan)

	for{
		_,ok := <-exitChan
		if !ok{
			break
		}
	}
}

