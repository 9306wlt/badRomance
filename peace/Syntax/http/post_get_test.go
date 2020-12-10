package http

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "testing"
)

/*
https://segmentfault.com/a/1190000013262746
 */
func http_post(){

    client := &http.Client{}

    req,err := http.NewRequest("post","baidu.com",strings.NewReader("name = cjb"))
    if err != nil{

    }
    req.Header.Set("content-type","application/x-www-form-urlencoded")
    req.Header.Set("cookie","name = anny")
    resp,err := client.Do(req)
    defer resp.Body.Close()
    body,err := ioutil.ReadAll(resp.Body)
    if err != nil{

    }
    fmt.Println(string(body))
}

func http_get(){

    resp,err := http.Get("https://www.baidu.com")
    if err != nil{
        panic(err)
    }
    defer resp.Body.Close()

    s,_ := ioutil.ReadAll(resp.Body)
    fmt.Printf(string(s))
}

func Test(t *testing.T){

    http_get()
}
