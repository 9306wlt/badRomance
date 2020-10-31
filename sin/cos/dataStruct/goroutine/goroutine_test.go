package goroutine

import (
	"strconv"
	"sync"
	"testing"
)

var wg = sync.WaitGroup{}

func Say(s string){
	println(s)
	//wg.Done 相当于 wg.Add(-1) 意思就是我这个协程执行完了
	wg.Done()
}

func Test_say(t *testing.T){

	////wg.Add(1)是说我有一个协程需要执行
	//wg.Add(1)
	//go Say("Hello World!")
	////wg.Wait() 就是告诉主线程要等一下，等协程都执行完再退出。
	//wg.Wait()
	//runtime.GOMAXPROCS(1)
	wg.Add(5)
	for i:=0;i<5;i++{
		go Say("Hello World:" + strconv.Itoa(i))
	}
	//for{
	//	time.Sleep(time.Second)
	//}
	wg.Wait()


}