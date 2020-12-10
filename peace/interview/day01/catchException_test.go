package day01

import (
	"fmt"
	"testing"
	"time"
)

/**
主要学习go语言中的异常 defer pannic recover
 */

func f(){
	defer func(){
		if r:=recover();r!=nil{
			fmt.Println("Recovered in f",r)
		}
	}()
	fmt.Println("Calling g.")

	//这里没有开辟新的协程，所以是在一个协程中（现在暂时的理解），所以发生了panic之后有recover,因此程序正常执行
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int){
	if i>3{
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v",i))
	}
	fmt.Println("Printing in g",i)
	//defer fmt.Println("Defer in g",i)

	g(i+1)
}


func a(){
	i :=0
	i++
	defer fmt.Println(i)
	i++
	return
}

func c() (i int){
	defer func() {
		i++
		fmt.Println(i)
	}()
	return 1
}

func G() {
	defer func() {
		//goroutine外进行recover
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("c")
	}()
	//创建goroutine调用F函数，这里是新建了一个协程，
	go F()
	time.Sleep(time.Second)
}

func F() {
	defer func() {
		fmt.Println("b")
	}()
	//goroutine内部抛出panic
	panic("a")
}

func TestException(t *testing.T){
	f()
	fmt.Println("Returned normally from f.")

	//value := c()
	//fmt.Println(value)

	G()
}
