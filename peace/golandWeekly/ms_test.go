package golandWeekly

import (
	"fmt"
	"testing"
)

func Test_(t *testing.T){
	a := [2]int{5, 6}
	b := [2]int{5, 6}
	if a == b {
		fmt.Println("equal")
		fmt.Printf("a的地址 =  %p\n",&a)
		fmt.Printf("b的地址 =  %p\n",&b)
	} else {
		fmt.Println("not equal")
	}

}
