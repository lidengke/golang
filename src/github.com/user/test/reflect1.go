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

  fmt.Println("============可写性 改变值============")
     var x float64 = 3.4;
    fmt.Println(x)
     r := reflect.ValueOf(&x)
     fmt.Println("type of r:", r.Type())
     fmt.Println("settability of r:", r.CanSet())

     v := r.Elem()
     fmt.Println("settability of v:", v.CanSet()) // true
     v.SetFloat(7.1)
     fmt.Println(v.Interface())  // 7.1
     fmt.Println(x)  // 7.1
}