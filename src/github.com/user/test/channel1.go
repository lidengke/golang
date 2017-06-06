package main

import "fmt"


// goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。那么goroutine之间如何进行数据的通信呢，Go提供了一个很好的通信机制channel。
// channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。这些值只能是特定的类型：channel类型。定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel：

// 	ci := make(chan int)
// 	cs := make(chan string)
// 	cf := make(chan interface{})
// channel通过操作符<-来接收和发送数据

//	ch <- v    // 发送v到channel ch.
//	v := <-ch  // 从ch中接收数据，并赋值给v

// 定义channel的时候 需要指明改channel传输的数据类型

func sum(a[] int, c chan int) {
    total := 0
    for _,v := range a {
        total += v
    }
    c <- total // send total to channel c, 想channel中写入数据将会阻塞当前调用，知道这个channel中的数据被读取。
}

func main () {
  a := []int{7,2,8,-9,4,0}
  c := make(chan int)
  go sum(a[:len(a)/2],c)
  go sum(a[len(a)/2:],c)

  x := <-c
  y := <-c

  fmt.Println(x,y,x+y)

//   默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。
//   所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（例如想ch中发送一个值 ch<-5）将会被阻塞，直到数据被读出。
//   无缓冲channel是在多个goroutine之间同步很棒的工具
 
}