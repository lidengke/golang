// 举个例子：fmt.Println是我们常用的一个函数,任何实现了String方法的类型都能作为参数被fmt.Println调用
package main

import (
    "fmt"
    "strconv"
)


type Stringer interface {
    String() string
}

type Human struct {
    name string
    age int
    phone string
}

func (h Human) String() string {
    return h.name + "-" + strconv.Itoa(h.age) + "-" + h.phone
}

func main() {
    bob := Human{"Bob",23,"000-777-xxxx"}
    fmt.Println(bob)
}

