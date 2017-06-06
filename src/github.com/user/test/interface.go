package main

import "fmt"

// 定义类型
type Human struct {
		name string
		age int
		phone string
}

type Student struct {
    Human //匿名字段Human
    school string
    loan float32
}

type Employee struct {
    Human //匿名字段Human
    company string
    money float32
}

// 定义接口
type Men interface {
    SayHi()
    Sing(lyrics string)
    Guzzle(beerstrein string)
}

type YoungChap interface {
    SayHi()
    Sing(song string)
    BorrowMoney(amount float32)
}

type ElderlyGent interface {
    SayHi()
    Sing(song string)
    SpendSalary(amount float32)
}

func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
    fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
    fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

func (s *Student) BorrowMoney(amount float32) {
    s.loan += amount
}

func (e * Employee) SpendSalary(amount float32) {
    e.money -= amount
}

func main () {
    stu := Student{Human{"mark",23,"13578155116"},"MIT",4125.01}
    employee := Employee{Human{"sam",31,"1546879213"},"google",1257.25}
    
    stu.SayHi();
    stu.Sing("hello,golang,you are the best program language")
    stu.Guzzle("nothing serious")
    fmt.Printf("%s has loan %f\n",stu.name,stu.loan)
    stu.BorrowMoney(1000.1)
    fmt.Printf("%s borrow  %f and now the total loan is %f\n",stu.name,1000.1,stu.loan)

     employee.SayHi();
    employee.Sing("hello,golang,you are the best program language")
    employee.Guzzle("nothing serious")
    fmt.Printf("%s has money %f\n",employee.name,employee.money)
    employee.SpendSalary(100)
    fmt.Printf("%s spend salary  %d and now the total loan is %f\n",employee.name,100,employee.money)

}


