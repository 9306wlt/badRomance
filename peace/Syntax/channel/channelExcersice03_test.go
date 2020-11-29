package channel

import (
	"fmt"
	"testing"
)

/*
需求：
统计1-8000中数中，哪些是素数

 */

func pushNum(intChan chan int){

	for i:=1;i<=8000;i++{
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int,prime chan int,exitChan chan bool){
	var flag bool
	for{
		num,ok := <-intChan
		if !ok{
			break
		}
		flag = true
		for i:=2;i<num;i++{
			if num%2 == 0{
				flag = false
				break
			}
		}
		if flag {
			prime <- num
		}

	}
	fmt.Printf("有一个协程因为取不到数据退出\n")
	exitChan <- true
}

func Test_channel03(t *testing.T){

	//存放数的管道
	intChan := make(chan int,1000)
	//开启一个协程去存放
	go pushNum(intChan)
	//存放素数的管道
	primeChan := make(chan int,2000)
	//标志退出的管道，当所有的协程都跑完，则这个管道中全为true
	exitChan := make(chan bool,4)

	//开启四个协程去计算8000个数中有哪些是素数
	for i:=0;i<4;i++{
		go primeNum(intChan,primeChan,exitChan)
	}

	//阻塞这里，直到所有的协程计算完毕才退出
	go func(){
		for i:=0;i<4;i++{
			<- exitChan
		}
		close(primeChan)
	}()

	for{
		num,ok := <- primeChan
		if !ok{
			break
		}
		fmt.Printf("素数为%d\n",num)
	}
	fmt.Printf("mian() 线程退出~")
}