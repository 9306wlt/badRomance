package channel

import (
	"fmt"
	"testing"
)

/**
需求：现在要求计算1-200各个数的阶乘，并且把结果放到map中
     最后显示出来，使用goroutine
*/

/**
思路：
	1、编写一个函数，用于计算各个数的阶乘，放到map中
	2、启动多个协程，分别结算结果放到map
	3、map应该是一个全局变量

*/

//改进版本 协程&通道.md  不同协程之间的通信---解决方案之一——全局变量加锁
var (
	myChannel = make(chan map[int]int,20)
	exitChan = make(chan bool,1)
)

//func test(myChannel chan map[int]int){
//	for i:=1;i<=20;i++{
//		myChannel <- Cal(i)
//	}
//	close(myChannel)
//}
//
//func Cal(n int)map[int]int{
//
//	res := 1
//	myMap := make(map[int]int,1)
//	for i:=1;i<=n;i++{
//		res *= i
//	}
//	myMap[n] = res
//	return myMap
//}


func calculate(n int){
	res := 1
	myMap := make(map[int]int,1)
	for i:=1;i<=n;i++{
		res *= i
	}
	myMap[n] = res
	myChannel <-  myMap
	if n == 20{
		exitChan <- true
		close(exitChan)
		close(myChannel)
	}
}

func read(){

}


func Test_test02(t *testing.T){

	//myChannel = make(chan map[int]int,20)
	for i:=1;i<=20;i++{
		go calculate(i)
	}
	for{
		_,ok := <-exitChan
		if !ok{
			break
		}
	}

	for{
		value,ok := <-myChannel
		if !ok{
			break
		}
		fmt.Printf("%v \n",value)
	}
}

