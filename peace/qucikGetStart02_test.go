package _interface

import (
	"fmt"
	"testing"
)

type AIntance interface {
	Say()
}

type Stu struct{
	Name string
}

func(stu Stu)Say(){
	fmt.Println("stu say")
}

func Test_interface(t *testing.T){
	var stu Stu
	stu.Say()
	var a AIntance = stu
	a.Say()
}
