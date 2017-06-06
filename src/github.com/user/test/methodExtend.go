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
}

type Employee struct {
    Human
    company string
}

func (h *Human) sayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}


func main() {
    mark := Student{Human{"mark",23,"125879112"},"MIT"}
    sam := Employee{Human{"sam",24,"158791111"},"Google"}
    mark.sayHi()
    sam.sayHi();
}