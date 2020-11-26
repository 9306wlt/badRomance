package _interface

import (
	"fmt"
	"testing"
)

//Monkey结构体
type Monkey struct {
	Name string
}

//声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (monkey *Monkey)Climbing(){
	fmt.Println(monkey.Name + "生来就会爬树")
}

type LittleMonkey struct {
	Monkey //继承
}

//让LittleMonkey实现BirdAble
func (littleMokey *LittleMonkey)Flying()  {
	fmt.Println(littleMokey.Name,"通过学习，会飞翔。。。")

}
//让LittleMonkey实现FishAble
func (littleMonkey *LittleMonkey)Swimming(){
	fmt.Println(littleMonkey.Name,"通过学习，会游泳。。。")
}



func Test_exercise02(t *testing.T){

	littleMokey := &LittleMonkey{Monkey{Name:"悟空"}}
	littleMokey.Climbing()
	littleMokey.Flying()
	littleMokey.Swimming()





}