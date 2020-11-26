package typeAssertions

import (
	"fmt"
	"testing"
)

func Test_assert(t *testing.T)  {
	var x interface{} //声明一个空接口
	var b2 float64
	x = b2 //空接口可以接受任何数据
	y := x.(float64) //类型断言
	fmt.Println(y)

}
