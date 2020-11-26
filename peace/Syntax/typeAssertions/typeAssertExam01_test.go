package typeAssertions

import (
	"fmt"
	"testing"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

func(p *Phone)Start(){
	fmt.Println(p.name+"手机开始工作")
}

func(p *Phone)Stop(){

	fmt.Println(p.name+"手机停止工作")
}

func (p *Phone)Call()  {
	fmt.Println(p.name+"手机正在通话中")

}

type Camera struct{
	name string
}

//让Phone实现Usb接口的方法
func (c *Camera)Start(){
	fmt.Println(c.name+"相机已通过usb连接")
}
func (c *Camera)Stop()  {
	fmt.Println(c.name+"相机断开Usb连接")
}


type Computer struct{

}

func(c *Computer)Working(usb Usb){
	usb.Start()
	//如果usb指向的是phone结构体，则还要调用Call方法
	//类型断言
	if phone,ok := usb.(*Phone);ok{
		phone.Call()
	}
	usb.Stop()
}

func Test_assertExam(t *testing.T){

	var usbArr [3]Usb
	usbArr[0] = &Phone{"vivo"}
	usbArr[1] = &Phone{"xiaomi"}
	usbArr[2] = &Camera{"尼康"}

	//遍历usbArr

	var computer Computer
	for _,v :=range usbArr{
		computer.Working(v)
		fmt.Println()
	}
}
