package method

import (
	"fmt"
	"testing"
)

/*
结构体上的简单方法例子
 */

type TwoInts struct {
	a int
	b int
}

func(tn *TwoInts) AddTem()int{
	return tn.a + tn.b
}

func(tn *TwoInts)AddToParam(param int)int{
	return tn.a + tn.b + param
}

func Test_method1(t *testing.T){
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("the sum is: %d\n",two1.AddTem())
	fmt.Printf("Add them to the param: %d\n",two1.AddToParam(20))

	two2 := TwoInts{3,4}
	fmt.Printf("the sum is:%d\n",two2.AddTem())

}