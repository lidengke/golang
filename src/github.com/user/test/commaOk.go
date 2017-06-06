package main
import(
    "fmt"
    "strconv"
)

// comma-ok 判别表达式

// 空类型
type Element interface{}

// 数组
type List [] Element

type Person struct {
    name string
    age int
}

//定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
    return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}

func main() {
    list :=make(List,3)
    list[0] = 1;
    list[1] = "hello,world"
    list[2] = Person{"mark",25}


    for index, element := range list {
        if value,ok := element.(int); ok {
            fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
        }else if value,ok := element.(string); ok {
             fmt.Printf("list[%d] is string and its value is %s\n", index, value)
        } else if value, ok := element.(Person); ok {
            fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
        } else {
            fmt.Printf("list[%d] is of a different type\n", index)
        }
    }


    fmt.Println("===========comma ok expression switch test=================")
    // `element.(type)`语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用`comma-ok`。
    for index,itemValue := range list {
        switch item := itemValue.(type) {
            case int:
                fmt.Printf("list[%d] is an int and its value is %d\n", index, item) 
            case string:
                fmt.Printf("list[%d] is string and its value is %s\n", index, item)
            case Person:
                fmt.Printf("list[%d] is a Person and its value is %s\n", index, item)   
            default :
                 fmt.Printf("list[%d] is of a different type\n", index)
        }
    }


}