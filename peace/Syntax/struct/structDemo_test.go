package _struct

import (
	"fmt"
	"peace/dataStruct/factory/model"
	"testing"
)

func Test_Demo(t *testing.T){
	var stu = model.NewStudent("jerry",12)
	fmt.Println(*stu)

}
