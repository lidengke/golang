package main

import "fmt"
import "math"
import "github.com/user/stringutil"
import "github.com/user/test"

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Println("Reverse: " + stringutil.Reverse("Hello.world"))
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))
	fmt.Printf("Pi is %g\n", math.Pi)
	var x = 1
	var y = 2
	fmt.Printf("%d + %d = %d\n", x, y, add(x, y))
	fmt.Printf("%d + %d = %d\n", x, y, add2(x, y))
	var a = "hello"
	var b = "world"
	a, b = swap(a, b)
	fmt.Printf("a = %s,b = %s\n", a, b)

	fmt.Println(split(17))

	forTest()
	forTest2()
	useForAsWhile()
	arrayTest()
	sliceTest()
	sliceTest2()
	pointerTest()
	var vertex test.Vertex
	vertex.X = 1
	vertex.Y = 2

	fmt.Println(vertex.X)
	fmt.Printf("x = %d, y = %d", vertex.X, vertex.Y)
}

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// return 没有显式返回值，此时将会默认返回 函数声明中定义的 变量 x y
	// 此时  x,y 初始化的值位0，通过一些函数内部的逻辑对x y进行变化
	return
}

func forTest() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("forTest sum = %d\n", sum)
}

func forTest2() {
	sum := 0
	i := 0
	for ; i < 10; i++ {
		sum += i
	}
	fmt.Printf("forTest2 sum = %d\n", sum)
}

// use for as while in C

func useForAsWhile() {
	sum := 0
	i := 0
	for i < 10 {
		i++
		sum += i
	}
	fmt.Printf("useForAsWhile sum = %d\n", sum)
}

func arrayTest() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}

func sliceTest() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] = %d\n", i, slice[i])
	}
}

func sliceTest2() {
	slice := []struct {
		a bool
		b int
	}{
		{true, 1},
		{false, 2},
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("index = %d,slice.a = %v,slice.b=%v\n", i, slice[i].a, slice[i].b)
	}
}

func pointerTest() {
	var p *int
	i := 43
	p = &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)
}
