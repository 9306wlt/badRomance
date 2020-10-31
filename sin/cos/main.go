package main

import (
	"sync"
	_"time"

	)

var wg = sync.WaitGroup{}

func Say(s string){
	println(s)
	wg.Done()
}

func main(){
	//wg.Add(1)是说我有一个协程需要执行
	wg.Add(1)
	go Say("Hello World!")
	wg.Wait()

}