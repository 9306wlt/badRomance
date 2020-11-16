package function

import (
	"fmt"
	"testing"
)

func f6(){
	fmt.Println("hello world")
}

func f7()int{
	return 10
}

//函数也可以作为参数的类型
func f8(x func()int){
	ret := x()
	fmt.Println(ret)
}

func f10(a,b int)int{
	return a+b
}

//函数还可以作为返回值
func f9(x func()int)func(int,int)int{
	return f10
}


func Test_funcType(t *testing.T){
	a := f6
	fmt.Println("%T\n",a) //函数f6的类型为 无参无返回值 func()
	b := f7
	fmt.Println("%T\n",b) //函数f7的类型为 无参有返回值 func()int

	f8(f7)
	f8(b)

	c := f9(f7)
	fmt.Println("%T\n",c)
}
