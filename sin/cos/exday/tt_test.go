package exday

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func getWeight(s string)float64{
	s = strings.ReplaceAll(s," ","")
	buf := []byte(s)
	var str string
	var num []int
	for i:=0;i<4;i++{
		num = append(num,int(buf[i] - 48))
	}
	for _,value := range num{
		str += strconv.Itoa(value)
	}
	weight,_ :=strconv.ParseFloat(str,64)
	return weight
}

func Test_getWeight(t *testing.T){
	var s string = " 0146    633     092"

	//a := "2031373534"
	//b,error := strconv.Atoi(a)
	//if error != nil{
	//	fmt.Println("字符串转换成整数失败")
	//}
	//fmt.Println(b)

	weight := getWeight(s)
	fmt.Println(weight)
}
