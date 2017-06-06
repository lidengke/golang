package  main

import "fmt"


type Human struct {
    name string
    age int 
    phone string
}
type Student struct {
    Human
    school string
    loan float32
}

type Employee struct {
    Human
    company string
    salary float32
}

func (h Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
    fmt.Println("La la, la la la, la la la la la...", lyrics)
}
func (h Student) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s and I am study at %s\n", h.name, h.phone,h.school)
}

func (h Employee) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s and I am work in %s\n", h.name, h.phone,h.company)
}

type Mem interface {
    SayHi()
    Sing(lyrics string)
} 

func main() {
    mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
    paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
    sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
    tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

    var i Mem
    i = mike

    fmt.Println("This is Mike, a Student:")
    i.SayHi()
    i.Sing("November rain")

    var p Mem
    p = paul

    fmt.Println("This is Paul, a Student:")
    p.SayHi()
    p.Sing("September")

    var s Mem
    s = sam

    fmt.Println("This is sam, an Employee:")
    s.SayHi()
    s.Sing("Born to be wild")


    //定义了slice Men
    fmt.Println("Let's use a slice of Men and see what happens")
    x := make([] Mem,3)
    x[0],x[1],x[2] = paul,sam,tom
    for _,h :=range x {
        h.SayHi()
    }

}

