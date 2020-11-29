package goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func Test_cpu(t *testing.T){
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)
}
