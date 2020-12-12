package 常见语法题二

import (
    "fmt"
    "testing"
)

type People struct {

}

func(p *People)ShowA(){
    fmt.Println(&p)
    fmt.Println("people showA")
    p.ShowB()
}

func(p *People)ShowB(){
    fmt.Println(&p)
    fmt.Println("people showB")

}

type Teacher struct {
    People
}

func(t *Teacher)ShowB(){
    fmt.Println("teacher showB")

}

func Test_question04(t *testing.T){
    teacher := Teacher{}
    //fmt.Printf("%p\n",&teacher)
    //teacher.ShowA()
    //teacher.People.ShowB()
    teacher.People.ShowB()


}
