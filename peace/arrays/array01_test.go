package arrays

import (
    "fmt"
    "testing"
)

func Test_array01(t *testing.T){
    a := [...]string{"a", "b", "c", "d"}
    for i := range a {
        fmt.Println("Array item", i, "is", a[i])
    }
}
