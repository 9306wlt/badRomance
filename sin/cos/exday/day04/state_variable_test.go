package day04

import (
	"fmt"
	"testing"
)

var(
	size = 1024
	max_size = size*2
)

func Test_stateVariable(t *testing.T){
	var a = 100
	fmt.Println(size,max_size,a)
}