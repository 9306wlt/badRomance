package applyAnonymous

import (
	"fmt"
	"testing"
)


func Adder2() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func Test_anonymous2(t *testing.T){

	var f = Adder2()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Println(f(300))
}
