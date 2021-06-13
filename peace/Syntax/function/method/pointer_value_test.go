package method

import (
	"fmt"
	"testing"
)

/*
如果想要方法改变接受者的数据，就在接受者的指针类型上定义定义该方法
 */

type B struct {
	thing int
}

func (b *B)change(){
	b.thing = 1
}

func(b B) write()string{
	return fmt.Sprint(b)
}

func Test_point_value(t *testing.T){
	var b1 B  //b1是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) //b2是指针
	b2.change()
	fmt.Println(b2.write())
}
