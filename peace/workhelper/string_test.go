package workhelper

import (
    "fmt"
    "testing"
)

func Test_string(t *testing.T){
    url := "https://paymgmt.shuzutech.com/invoice/h5/selfService?appId="
    appId := "hsrtjtydjftukfhkhjkj"
    version := "1.0.1"
    //appSecret := "e6169f9e95cf4446c21c11d4e4c611100ab2473d"
    encryptMsg := "sfhksdjfiljsdalkd"
    a := url+appId+"&encryptMsg="+encryptMsg+"&version="+version
    fmt.Println(a)
}
