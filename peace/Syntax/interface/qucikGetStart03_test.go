package _interface

import (
	"fmt"
	"testing"
)

type BIntance interface {
	test01()
}

type CIntance interface {
	test02()

}

type DIntance interface {
	BIntance
	CIntance
	test03()

}


//如果要实现DInterface，就需要将BInstance CInstance的方法都实现
type Student struct{

}

func(stu Student)test01(){
	fmt.Println("stu say  BInstance")
}

func(stu Student)test02(){
	fmt.Println("stu say  CInstance")
}

func(stu Student)test03(){
	fmt.Println("stu say  DInstance")
}

func Test_interface03(t *testing.T){
	var stu Student

	var b BIntance = stu
	b.test01()
}
