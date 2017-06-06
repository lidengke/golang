package main

import (
    "fmt"
)


func read(c chan int) {
     fmt.Println("blocked by channel,wait to read value from channel")
     v := <- c
     fmt.Println("read value from channel:",v)
}

func main() {
    c := make(chan int ,2) // 创建一个缓冲区为2 int类型的缓冲channel

    fmt.Println("send 1 to channel without blocking")
    c <- 1
    fmt.Println("send 2 to channel without blocking")
    c <- 2

    fmt.Println("sending 3 to channel will block until another goroutine read value from this channel")
    go read(c)
    c <-3
    fmt.Println("another goroutine read value from this channel, program exit success !")

}


