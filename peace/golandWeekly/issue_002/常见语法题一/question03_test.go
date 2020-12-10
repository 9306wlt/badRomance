package 常见语法题一

import (
    "encoding/json"
    "fmt"
    "testing"
)

type People struct{
    name string `json:"name"` //私有属性不应该加json的标签
}

/*
type People struct{
    Name string `json:"name"` 把name改成首字母大写就可以运行并且有值，私有属性在json解码或者转码的时候无法上线私有属性的转换，换句话说 私有变量是不可导出的
}
 */

func Test_question03(t *testing.T){
    js :=`{
            "name":"11"
         }`
    var p People
    err := json.Unmarshal([]byte(js),&p)
    if err != nil{
        fmt.Println("err:",err)
        return
    }
    fmt.Println("people:",p)
}

