package day04

import (
	"fmt"
	"testing"
)

/**

先看下append函数原型，func append(slice []Type, elems ...Type) []Type
1、第一个参数是一个Type类型的切片，第二个参数是可变参数，类型为Type。
2、在go中可变参数的工作原理是把可变参数转换为一个新的切片。
3、如果将一个切片传给一个可变参数，这种情况无法编译，原因很直接，由可变参数定义可知，elems ...Type意味着它可以
接收Type类型可变参数。s1 = append(s1, s2)，这里，s2已经是一个int类型的切片，编译器试图在s2的基础上再创建一个切片，所以会失败，因为s2是一个[]int类型，而不是int类型。
4、或者我们在s1 = append(s1, s2...),就可以通过编译了，如果是这种写法，切片将直接传入参数，不再创建新的切片


 */

func Append_test(){
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1,cap(s1))
}

func Test_Append(t *testing.T){
	Append_test()
}
