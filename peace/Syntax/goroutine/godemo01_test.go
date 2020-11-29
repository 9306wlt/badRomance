package goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/**
 在主线程中开启一个协程，每隔一秒输出“hello world”
 主线程中也每隔一秒输出hello golang,输出10此后退出
 要求主线程和goroutine同时进行
 */

func test(){
	for i:=1;i<11;i++{
		fmt.Println("test() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func Test_test(t *testing.T){

	go test()
	for i:=1;i<11;i++{
		fmt.Println("main() hello golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
