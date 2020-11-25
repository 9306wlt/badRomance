package _interface

import (
	"fmt"
	"testing"
)

//声明定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct{

}

//让Phone实现Usb接口的方法
func (p Phone)Start(){
	fmt.Println("手机已通过usb连接")
}
func (p Phone)Stop()  {
	fmt.Println("手机断开Usb连接")
}


type Camera struct{

}

//让Phone实现Usb接口的方法
func (c Camera)Start(){
	fmt.Println("相机已通过usb连接")
}
func (c Camera)Stop()  {
	fmt.Println("相机断开Usb连接")
}


type Computer struct{

}

//编写一个方法working方法，接收一个Usb接口类型变量
//只是要实现了Usb接口（所谓实现usb接口，就是指实现了Usb接口声明所有方法）

func (c Computer)Working(usb Usb)  {//usb变量会根据传入的实参，来判断到底是Phone还是Camera

	//通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func Test_Interface(t *testing.T)  {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)

}

