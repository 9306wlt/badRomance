package channel

import (
    "fmt"
    "testing"
)

/*

对于同一通道，发送操作在接受者准备好之前是不会结束的。这就意味着，如果一个无缓冲通道在没有空间接收数据的时候，新的输入数据无法输入，即发送者处于阻塞状态。
对于同一通道，接收操作是阻塞的，直到发送者可用。如果通道中没有数据，接收者会保持阻塞。

以上两条性质，反映了无缓冲通道的特性：同一时间只允许至多一个数据存在于通道中。

作者：_Hotown
链接：https://juejin.cn/post/6844903778617917453
来源：掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

func pump04(ch chan int){
    for i:=0; ;i++{
        ch <- i
    }
}

func Test_gc04(t *testing.T) {
    ch1 := make(chan int)
    go pump04(ch1)
    fmt.Println(<-ch1)
}
