package main

import (
	"runtime"
	"fmt"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        // runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
        runtime.Gosched() 
        fmt.Println(i,"-",s)
    }
}

func main() {
    var s1,s2 string = "hello","goroutine"
    go say(s2)
    say(s1)
}

