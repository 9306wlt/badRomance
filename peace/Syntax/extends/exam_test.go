package extends

import (
	"fmt"
	"testing"
)

//编写一个学生考试系统

//小学生
type Pupil struct {
	Name string
	Age int
	Score float64
}

//现实他的成绩
func (pupil *Pupil) ShowInfo(){
	fmt.Printf("姓名 = %v ,年龄 = %v ,得分 = %v\n",pupil.Name,pupil.Age,pupil.Score)
}

func (pupil *Pupil) SetScore(score float64){
	pupil.Score = score
}

func (pupil *Pupil) testing(){
	fmt.Println("正在考试中")
}

type Graduate struct {
	Name string
	Age int
	Score float64
}

//现实他的成绩
func (graduate *Graduate) ShowInfo(){
	fmt.Printf("姓名 = %v ,年龄 = %v ,得分 = %v\n",graduate.Name,graduate.Age,graduate.Score)
}

func (graduate *Graduate) SetScore(score float64){
	graduate.Score = score
}

func (graduate *Graduate) testing(){
	fmt.Println("正在考试中")
}

//代码冗余

func Test_exam(t *testing.T){
	//测试
	var p = &Pupil{
		Name:  "tom",
		Age:   10,
		Score: 0,
	}
	p.testing()
	p.SetScore(90)
	p.ShowInfo()

	var g = &Graduate{
		Name:  "mary",
		Age:   20,
		Score: 0,
	}
	g.testing()
	g.SetScore(90)
	g.ShowInfo()

}

