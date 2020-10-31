package day02

import (
	"fmt"
	"testing"
)

func testForRange(){
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key,val := range slice {
		m[key] = &val
	}

	for k,v := range m {
		fmt.Println(k,"->",*v)
	}
}

func Test_testForRange(t *testing.T){
	testForRange()
}
