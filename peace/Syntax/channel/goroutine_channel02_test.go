package channel

import (
	"fmt"
	"testing"
	"time"
)

func sendLove02() {
	fmt.Println("Love you,mm...")
}

func responseLove02() {
	time.Sleep(6 * 1e9)
	fmt.Println("Love you, too")
}

func waitFor02() {
	for i := 0; i < 5; i++ {
		fmt.Println("Keep waiting...")
		time.Sleep(1 * 1e9)
	}
}

func Test_gc02(t *testing.T) {

	fmt.Println("Reveal romantic feelings....")
	go sendLove02()
	go responseLove02()
	waitFor02()
	fmt.Println("Leaving 含泪")
}
