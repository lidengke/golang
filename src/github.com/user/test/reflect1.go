package main
import (
    "fmt"
    "reflect"
)

func main() {
    type MyInt int
    var i MyInt = 1
    var mtype reflect.Type = reflect.TypeOf(i)
    fmt.Println(mtype)
    fmt.Println(mtype.Kind())

    fmt.Println("===========================")
    var mValue reflect.Value = reflect.ValueOf(i)
    fmt.Println(mValue)
    fmt.Println(mValue.Kind())
    fmt.Println(mValue.Int())
    //fmt.Println(mValue.Set(10))


     fmt.Println("============可写性============")

     canset := mValue.CanSet()
       fmt.Println(canset)
}