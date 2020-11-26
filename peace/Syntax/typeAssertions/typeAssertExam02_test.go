package typeAssertions

import (
	"fmt"
	"testing"
)

func TypeJudge(items... interface{}){
	for index,x := range items{
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n",index,x)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n",index,x)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n",index,x)
		case int,int32,int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n",index,x)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n",index,x)
		case Student:
			fmt.Printf("第%v个参数是Student类型，值是%v\n",index,x)
		case *Student:
			fmt.Printf("第%v个参数是*Student类型，值是%v\n",index,x)
		default:
			fmt.Printf("第%v个参数是 类型不确定 ，值是%v\n",index,x)

		}
	}
}

type Student struct {
	Name string
	Age int

}

func Test_assertExam02(t *testing.T){

	var n1 float32 = 1.1
	var n2 float64 = 1.2
	var n3 int32 = 3
	var name string = "tony"
	address := "深圳"
	n4 := 300
	var stu1 = Student{
		Name:  "smith",
		Age:   10,

	}
	var stu2 = &Student{
		Name:  "smith",
		Age:   10,

	}
	TypeJudge(n1,n2,n3,name,address,n4,stu1,stu2)
}
