package _interface

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

//实现对Hero结构体切片的排序

//声明Hero结构体
type Hero struct {
	Name string
	Age int
}

//声明一个结构体切片类型
type HeroSlice []Hero


//实现Interface接口
func (hs HeroSlice)Len() int  {
	return len(hs)
}

//Less方法就是决定你使用什么标准进行排序
//1、按Hero的年龄从小到大排序

func(hs HeroSlice)Less(i,j int)bool{
	//return hs[i].Age < hs[j].Age

	//修改成对Name排序
	return hs[i].Name <hs[j].Name
}

func(hs HeroSlice)Swap(i,j int){
	//交换
	hs[i],hs[j] = hs[j],hs[i]
}

type Students struct {
	Name string
	Age int
	Score float64
}


func Test_exercise(t *testing.T){
	//将Students的切片，按Score从大到小排序

	//先定义一个数组/切片
	var intSlice = []int{0,-1,10,7,90}

	//要求对intSlice切片进行排序
	//使用排序算法进行排序
	//也可以使用系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对结构体切片进行排序
	var heros HeroSlice
	for i:=0;i<10;i++{
		hero := Hero{
			Name: fmt.Sprintf("英雄%d",rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//将hero append到 heros切片
		heros = append(heros,hero)
	}
	//调用前
	for _,v := range heros{
		fmt.Println(v)
	}
	//调用sort.Sort
	sort.Sort(heros)
	fmt.Println("----------排序后-------")
	for _,v := range heros{
		fmt.Println(v)
	}
}