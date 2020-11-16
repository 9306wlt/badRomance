package function

import "testing"

func f(x int,y int)(sum int){
	sum  = x + y //使用命名的返回值，那么在函数中可以直接使用返回值变量
	return       //使用命名的返回值，return后面可以省略返回值变量

}

func Test_function(t *testing.T){
	f(100,200)
}
