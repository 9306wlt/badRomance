package day04

import (
	"fmt"
	"testing"
)

func Point_Address(){
	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"

	//取字符串的地址
	//fmt.Printf("house address: %p\n", &house)
	//fmt.Println(&house)
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	fmt.Println(&house)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(&ptr)
	// 打印ptr的类型
	//fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	//fmt.Printf("address  1: %v\n", ptr)
	fmt.Println("wojiuxiangkankan",*ptr)
	fmt.Printf("address  2: %p\n", &ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

}

func Test_new(t *testing.T){

	//Point_Address()

	var list *[]int = new([]int)
	//输出list类型,从结果来看其类型为*[]int 表示指向数组的一个指针
	fmt.Printf("%T\n",list)

	//对指针进行取值操作，取出地址指向的值，因为list指向[]int，它存的也就是[]int的地址，取得也就是[]int的值
	fmt.Println("取出指针所指向变量的值 []int value = ：",*list)
	fmt.Printf("*list 的 Type = %T\n：",*list)

	//&表示取出地址，取出的是指针变量list的地址
	fmt.Println("ptr address = ：",&list)

    //输出指针list的值,也就是[]int的地址
	fmt.Printf("[]int address = %p\n",list)

	//输出list自身的地址
	fmt.Printf("ptr address : %p\n",&list)

	*list = append(*list,1,2,3,4)
	fmt.Println("取出指针list所指向变量的值 []int value = ：",*list)
}
