package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
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

//改进版本 加锁  不同协程之间的通信---解决方案之一——全局变量加锁
var(
	lock sync.Mutex


	myMap03 = make(map[int]int,10)
)

func goroutinetest03(n int){

	res := 1
	for i:=1;i<=n;i++{
		res *= i
	}

	//加锁
	lock.Lock()
	myMap03[n] = res
	//解锁
	lock.Unlock()

}

func Test_test03(t *testing.T){

	//concurrent map writes
	for i:=1;i<=20;i++{
		go goroutinetest03(i)
	}

	time.Sleep(time.Second*5)

	for index,value := range myMap03{
		fmt.Printf("myMap[%d] = %d\n",index,value)
	}
}
