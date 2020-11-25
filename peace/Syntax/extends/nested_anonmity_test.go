package extends

import (
	"fmt"
	"testing"
)

//嵌套匿名结构体

//编写一个学生考试系统
type Student struct {
	Name string
	age int
	Score float64
}

func(stu *Student)ShowInfo(){
	fmt.Printf("姓名 = %v ,年龄 = %v ,得分 = %v\n",stu.Name,stu.age,stu.Score)
}

func(stu * Student)SetScore(score float64){
	//其他业务 校验分数合法性
	stu.Score = score
}

func(stu *Student)testing(){
	fmt.Println("正在考试")
}

type Pupil2 struct {
	Student //嵌入了Student匿名结构体
}

func(p * Pupil2)testing(){
	fmt.Println("小学生正在考试")
}

type Graduate2 struct {
	Student  //嵌入了Student匿名结构体
}

func (g *Graduate2)testing(){
	fmt.Println("大学生正在考试")
}

func Test_nested_anonmity(t *testing.T){

	//
	var p  = &Pupil2{}
	p.Student.Name = "tom~"
	p.Student.age = 8
	p.testing()
	p.Student.SetScore(70)
	p.Student.ShowInfo()

	var g  = &Graduate2{Student{
		Name:  "smith",
		age:   23,

	}}
	g.testing()
	g.Student.SetScore(89)
	g.Student.ShowInfo()
}