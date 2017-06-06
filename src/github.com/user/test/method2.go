package main


import "fmt"


const(
    WHITE = iota
    BLACK
    BLUE
    RED
    YELLOW
)

type Color byte

// box 定义一个结构体
type box struct {
    width,height,depth float64
    color Color
}

type BoxList []box // a slice of box box的一个切片


func (c *Color) modifyColor(c2 Color) {
      *c  = c2; 
}


func (b box) Volume() float64 {
    return b.width * b.height * b.depth
}

// SetColor set color
func(b *box) SetColor(c Color) {
    b.color = c // 等价于 *box.color = c //GO自动识别
}
// BiggestColor return the biggest color
func (boxList BoxList) BiggestColor() Color {
    v :=0.00
    k :=Color(WHITE)
    for _,b :=range boxList {
        if bv := b.Volume(); bv > v {
            v = b.Volume()
            k = b.color
        }
    }
    return k
}

// 也就是说：

// 如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method
// 类似的

// 如果一个method的receiver是T，你可以在一个*T类型的变量P上面调用这个method，而不需要 *P去调用这个method

func (bl BoxList) PlainItBlack() {
    for i, _:= range bl {
        bl[i].SetColor(BLACK) //
    }
}

func  (c Color) String() string {
    strings :=[]string{"WHITE","BLACK", "BLUE", "RED", "YELLOW"}
    return strings[c]
}


func main(){
    boxes := BoxList {
        box{4,4,4,RED},
        box{10, 10, 1, YELLOW},
			box{1, 1, 20, BLACK},
			box{10, 10, 1, BLUE},
			box{10, 30, 1, WHITE},
			box{20, 20, 20, YELLOW},
    }

    fmt.Printf("We have %d box in boxed\n",len(boxes));
    fmt.Println("The column of the first on is",boxes[0].Volume(),"cm³")
    fmt.Println("The color of the last one is ",boxes[len(boxes) -1].color.String())
    fmt.Println("The biggest one is", boxes.BiggestColor().String())

    fmt.Println("Let's paint them all black")
    boxes.PlainItBlack()
    fmt.Println("The color of the last one is ",boxes[len(boxes) -1].color.String())
    fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())


     cl := Color(BLACK)
     cl.modifyColor(Color(YELLOW))
     fmt.Println("The color is ",cl)
    (&cl).modifyColor(Color(RED))
    fmt.Println("The color is ",cl)

}
