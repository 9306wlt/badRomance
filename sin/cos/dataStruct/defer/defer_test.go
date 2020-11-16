package _defer

import (
	"fmt"
	"testing"
)

func deferDemo(){
	fmt.Println("start")
	defer fmt.Println("running") //defer把它后面的语句延迟到函数即将返回的时候再执行
	fmt.Println("end")
}

//go语言中函数的return不是原子操作，在底层是分两步来执行
//第一步：返回值赋值
//第二步：真正的RET返回
//函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

func f1() int{
	x := 5
	defer func(){
		x++
	}()
	return x //把x的值拷贝一份作为返回值
}
//f1()返回的是5，因为是无命名值的返回，后面defer语句修改的是x的值，不是返回值

func f2()(x int){
	defer func() {
		x++
	}()
	return 5
}
//f2()返回的是6，因为返回值是命名的，所以5赋值给了x，后面defer又对返回值x做了修改，所以返回的是6

func f3()(y int){
	x := 5
	defer func() {
		x++
	}()
	return x
}
//f3()返回的是5，因为返回值y = x = 5,后面修改的是x的值

func f4()(x int){
	defer func(x int) {
		x++ //改变的是函数中的副本
	}(x)
	return 5
}
//f4()返回的是5，因为返回值 x = 5，后面又将返回值作为defer函数的参数，defer函数中声明了x变量

func Test_defer(t *testing.T){
	deferDemo()
}