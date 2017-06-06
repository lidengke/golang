package main

import "fmt"

func main(){
    var a interface{} // 定义一个空接口
    var i int = 5;
    a = i;

    fmt.Println(a)
    var s = "hello,world"
    a = s
    fmt.Println(a)
}