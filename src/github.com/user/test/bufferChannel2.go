package main
import "fmt"


func fibonacci(n int,c chan int) {
    x,y := 1,1
    for i := 0; i < n; i++ {
        c <- x
        t := y
        y = x + y
        x = t;
        // 很奇怪的表达式，这种写法居然和上面 三行代码的效果一直, 并不是常规理解上的先将y赋值给x,然后再执行 y = x+y
        // 反而更像是 x = y 然后在执行y=x+y时，x的值仍然是未执行x=y赋值操作之前的值。 
        // x,y = y,x+y 
        fmt.Printf("x = %d,y = %d\n",x,y)    
    }
    close(c)
}


func main(){
    c := make(chan int,5)
    go fibonacci(cap(c),c)
    for i := range c {
        fmt.Println(i)
    }
}

// for i := range c 能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显式的关闭channel，生产者通过内置函数close关闭channel。
// 关闭channel之后就无法再发送任何数据了，在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。

// 记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
// 另外记住一点的就是channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的

