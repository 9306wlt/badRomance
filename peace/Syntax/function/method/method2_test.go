package method

import (
	"fmt"
	"testing"
)

/*
非结构体类型上方法的例子
 */


type InVector []int

func(v InVector) Sum()(s int){
	for _,x := range v{
		s += x
	}
	return
}

func Test_Method2(t *testing.T){
	fmt.Println(InVector{1,2,3}.Sum())
}