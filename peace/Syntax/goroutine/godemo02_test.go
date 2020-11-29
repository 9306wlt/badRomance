package goroutine

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

var(
	myMap = make(map[int]int,10)
)

func goroutinetest(n int){

	res := 1
	for i:=1;i<=n;i++{
		res *= i
	}

	myMap[n] = res

}

func Test_test02(t *testing.T){

	//concurrent map writes
	for i:=1;i<=200;i++{
		go goroutinetest(i)
	}

	for index,value := range myMap{
		fmt.Printf("myMap[%d] = %d\n",index,value)
	}
}
