package channel

import (
	"fmt"
	"testing"
)

type Cat struct {
	Name string
	Age int
}

func Test_channel(t *testing.T){
	var allchan chan interface{}
	allchan = make(chan interface{},5)
	allchan <- 10
	allchan <- "tom"
	cat := Cat{
		Name: "tom~",
		Age: 2,
	}
	allchan <- cat

	//想要取得第三个元素，那么就得先把前两个元素从管道推出来
	<-allchan
	<-allchan
	newCat := <-allchan

	fmt.Printf("newCat type = %T  value = %v \n",newCat,newCat)

	//下面的用法是错误的，编译不通过
	//fmt.Printf("newCat.Name = %v\n",newCat.Name)

	//解决方法是：使用类型断言，因为管道中存放的类型是interface
	a := newCat.(Cat)
	fmt.Printf("newCat.Name = %v\n",a.Name)
}
