GO语言是静态类型的语言，和C++/Java一样，变量声明后在编译期其静态类型就确定了。

var i int
type Myint int
var j Myint

1. i 为 int类型, j为Myint类型, i 和 j 是不同的静态类型,虽然二者具备相同的底层数据类型。他们不能直接相互赋值，需要进行转换。
2. interface是一个十分特殊的类型,它是一批方法的集合。一个接口类型的变量可以存储任何类型的详细真实的值，只要这个值的类型实现了这个接口方法。
3. 一个接口类型的变量,无论最终存储的真实值是何种类型,但其静态类型依然是声明时指定的类型。
案例如下：
// Reader is the interface that wraps the basic Read method.
type Reader interface {
    Read(p []byte) (n int, err error)
}

任何一个类型实现了上述这个签名的Read方法,就可以称之为实现了Reader接口。对于一个Reader接口类型的变量可以存储任何实现了Reader的值。
var r Reader

r = os.Stdin  // 标准输入
r = bufio.newBuffer() // buffer io
r = new(bytes.Buffer) // 数组buffer

上例说明中: r 存储三种不同Reader的实现类型的值。虽然r存储Reader三种不同实现类型的值，但无论r被赋予何种值，其静态类型仍然是Reader

接口和反射之间由很大的关联，先看看接口变量的表示。


接口变量的表示

1. 一个接口变量存储了一对值：
    (1)赋予这个变量的真实具体值;
    (2)这个具体真实值的类型描述。
   更加准确的说,真实具体的值就是底层数据值,而类型描述则是对这个底层数据值的类型描述。

// The static type of the interface determines what methods may be invoked with an interface variable, even though the concrete value inside may have a larger set of methods.
2. 一个变量静态类型决定了这个变量能够调用哪些方法,即使真实数据值的底层数据类型包含更多的方法,但只能调用这个静态类型包含的方法。


举例说明:
var r io.Reader
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
if err != nil {
    return nil, err
}
r = tty

变量r包含了一对值(tty,*os.File),其中存储的底层数据是tty,底层数据类型是*os.File。需要注意的是: *os.File 实现了Read方法，并且还包含了更多其他的方法。
虽然接口变量r只提供Read方法，但是对与真实值tty而言，还包含了Write方法以及其他更多的方法。所以，可以通过断言方式进行如下的转型:
var w io.Writer
w = r.(io.Writer) 
这是一个断言表达式,它断定变量r内部的实际值也实现了io.Writer接口。所以才能赋值给w。 赋值之后 w就指向了(tty,*os.File) 这是 w和r包含了相同的值对(tty,*os.File)


One important detail is that the pair inside an interface always has the form (value, concrete type) and cannot have the form (value, interface type). Interfaces do not hold interface values.
重要的一点: 接口变量包含的值对通常是有(底层数据值,底层数据类型)的方式组成，而不是(底层数据，接口类型)组成。 接口类型不能存储接口值。 接口类型是一种方法集合的描述，需要要具体实现。 
值对中的底层数据类型必须真实的了类型（基本类型或者struct）


下面来看看反射定律:

** 第一条反射定律 ： 从接口变量到反射对象

反射是一个检测存储在接口变量中(值对)的机制。在开始之前，我们需要知道Reflection包中两个类型：Type 和 Value。 这两个类型提供一种访问这个变量内容的方式。
两个简单的函数调用： 调用reflect.TypeOf返回reflect.Type, reflect.ValueOf返回reflect.Value,分别对应这个接口变量的reflect.Type 和 reflect.Value。 
当然通过reflect.Value是很容易得到reflect.Type的,这里先将其分开。

reflect.Value 和 reflect.Type包含有很多的方法，下面列举部分:
1. reflect.Value 有一个Type()方法返回reflect.Type
2. reflect.Value和reflect.Type都具有一个Kind()方法，返回一个常量，表示这个变量存储的数据的真实底层数据类型。常见值有：Uint、Float64、Slice等
3. reflect.Value还有一些Int() Float() 方法，用来提取这个变量的底层数据，Int()提取int64 Float()提取float64. 
4. 还有一些用来修改数据的方法，比如SetInt、SetFloat，在讨论它们之前，我们要先理解“可修改性”（settability），这一特性会在“反射第三定律”中进行详细说明。
反射库提供了很多值得列出来单独讨论的属性。首先是介绍下Value 的 getter 和 setter 方法。为了保证API 的精简，这两个方法操作的是某一组类型范围最大的那个。
比如，处理任何含符号整型数，都使用 int64。也就是说 Value 类型的Int 方法返回值为 int64类型，SetInt 方法接收的参数类型也是 int64 类型。实际使用时，可能需要转化为实际的类型：
var x uint8 = 'x'
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())       // uint8.
fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
x = uint8(v.Uint())    // v.Uint returns a uint64.
第二个属性是反射类型变量（reflection object）的 Kind 方法 会返回底层数据的类型，而不是静态类型。如果一个反射类型对象包含一个用户定义的整型数，看代码：
type MyInt int
var x MyInt = 7
v := reflect.ValueOf(x)
上面的代码中，虽然变量 v 的静态类型是MyInt，不是 int，Kind 方法仍然返回 reflect.Int。换句话说， Kind 方法不会像 Type 方法一样区分 MyInt 和 int。

TypeOf() 返回的是静态类型,Kind() 返回的是底层数据类型

** 第二条反射定律 ： 从反射对象到接口变量

1. 和物理学中的反射一样，通过反射也能生成原始对照物。
2. 对于一个已知的reflect.Value，我们可以通过Interface()方法恢复出一个接口类型变量。事实上，这些方法将类型和值信息组装成一个接口的描述后返回。
reflect.Value有一个Interface()的方法，如下所示:
func (value reflect.Value) Interface() interface{}
有了Value的这个方法，然后我们可以通过断言，恢复底层的真实的值。
    y := v.Interface().(float64) //断言v的真实值为float64 并赋值给y

上面这段代码会打印出一个 float64 类型的值，也就是 反射类型变量 v 所代表的值。
事实上，我们可以更好地利用这一特性。标准库中的 fmt.Println 和 fmt.Printf 等函数都接收空接口变量作为参数，fmt 包内部会对接口变量进行拆包（前面的例子中，我们也做过类似的操作）。
因此，fmt 包的打印函数在打印 reflect.Value 类型变量的数据时，只需要把 Interface 方法的结果传给 格式化打印程序：
fmt.Println(v.Interface())
你可能会问：问什么不直接打印 v ，比如 fmt.Println(v)？ 答案是 v 的类型是 reflect.Value，我们需要的是它存储的具体值。由于底层的值是一个 float64，我们可以格式化打印：
fmt.Printf("value is %7.1e\n", v.Interface())
上面代码的打印结果是：
3.4e+00
同样，这次也不需要对 v.Interface() 的结果进行类型断言。空接口值内部包含了具体值的类型信息，Printf 函数会恢复类型信息。
简单来说，Interface 方法和 ValueOf 函数作用恰好相反，唯一一点是，返回值的静态类型是 interface{}。


** 第三条反射定律 ： 修改反射对象的前提是接口变量可以被修改
这条规则很容易让人混淆，但是如果理解了第一条规则，则比较容易理解。下面的代码不能正确运行的：
var i float64 = 3.4
 v := reflect.ValueOf(i) // 由接口变量获取反射对象
 v.SetFloat(7.1) // error will panic
 执行这段代码，会引起一个panic: reflect.Value.SetFloat using unaddressable value

 引起这个恐慌的原因并不是因为值7.1不能被寻址的,而是因为v不可写。 "可写性"是 reflect.Value的一个属性,并不是所有的反射类型变量都可写。
 CanSet()这个方法可以判别Value是否可行性，例如:
 var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("settability of v:", v.CanSet()) -->settability of v: false

可写性有点性可寻址性，但是更为严格。他是反射类型变量R的一个属性，可以修改创建这个反射类型变量R的原始接口类型变量V。它是反射类型变量的一种属性，赋予该变量修改底层存储数据的能力。
反射类型变量的"可写性"最终是由一个事实决定的：反射对象是否存储了原始值。举个代码例子：
var x float64 = 3.4
v := reflect.ValueOf(x)

执行reflect.ValueOf(x)这行代码，实际上是将变量x的一个copy作为实际参数传递给了reflect.ValueOf()方法，而不是将变量x本身传给reflect.ValueOf()。
所以返回的反射类型变量v中并没有保存x本身。而只是保存了x的一个copy变量。

我们假设： v.SetFloat(7.1)是允许的并且执行成功的，这个操作也不会更新变量x，尽管看起来v是从变量x创建来的。相反，这个操作会更新传入的实参变量x变量的copy，
并且这个x的copy的变量会存储在返回的反射对象中，而对于变量x而言，则没有任何影响，那么即使这个SetFloat(7.1)的操作时允许的，但实际上也是无用的让人困惑的。
所以，干脆就限制这种徒劳的操作为非法的，而可行性就是为了避免这种问题的存在。

这看上去很诡异，但是事实上很常见，在函数的传递过程中就会经常出这种情况，例如：

传递一个参数x 给函数 f(x) 函数f改变x的值，因为传递给f的是x的一个copy，而不是x自身(c语言中的值传递和指针传递的区别)。 如果想通过f改变x的值，那么就需要将变量x
的地址传给函数f(也就是指针)。 

反射的工作原理和上述的场景是一样的，如果希望通过反射来改变变量x的值，那么就需要给reflect.ValueOf()传递当前变量x的地址(指针)

var  x float64 = 3.4

r := reflect.ValueOf(&x) // 传递给ValueOf()的是一个地址，而不是x的变量copy

fmt.Println("type of r:", r.Type())
fmt.Println("settability of r:", r.CanSet())

输出如下:
type of r: *float64
settability of r: false

r是一个反射对象,r绑定的静态类型为*float64
r.CanSet() 返回false。表明r是不可写的 ? 上文中不是说可以通过传递地址来修改变量的值么,这里为什么返回的确是不可写,怎么回事呢?

需要说明的是：
r是一个反射对象变量，他type绑定的是一个指针。 实际上我们也不是要修改r，而是要修改p指向的变量。为了得到 r指向的数据，
可以调用 Value 类型的 Elem 方法。Elem 方法能够对指针进行“解引用”，然后将结果存储到反射 Value类型对象 v中：

v := r.Elem()
fmt.Println("settability of v:", v.CanSet()) // true
v.SetFloat(7.1)
fmt.Println(v.Interface())  // 7.1
fmt.Println(x)  // 7.1


结构体（struct）
在前面的例子中，变量 v 本身并不是指针，它只是从指针衍生而来。把反射应用到结构体时，常用的方式是 使用反射修改一个结构体的某些字段。只要拥有结构体的地址，我们就可以修改它的字段。
下面通过一个简单的例子对结构体类型变量 t 进行分析。
首先，我们创建了反射类型对象，它包含一个结构体的指针，因为后续会修改。
然后，我们设置 typeOfT 为它的类型，并遍历所有的字段。
注意：我们从 struct 类型提取出每个字段的名字，但是每个字段本身也是常规的 reflect.Value 对象。
type T struct {
 A int
 B string
}
t := T{23, "skidoo"}
s := reflect.ValueOf(&t).Elem()
typeOfT := s.Type()
for i := 0; i < s.NumField(); i++ {
 f := s.Field(i)
 fmt.Printf("%d: %s %s = %v\n", i,
  typeOfT.Field(i).Name, f.Type(), f.Interface())
}
上面这段代码的输出如下：
A int = 23
B string = skidoo
这里还有一点需要指出：变量 T 的字段都是首字母大写的（暴露到外部），因为struct中只有暴露到外部的字段才是“可写的”。
由于变量 s 包含一个“可写的”反射对象，我们可以修改结构体的字段：

f.Interface())s.Field(0).SetInt(77)
s.Field(1).SetString("Sunset Strip")
fmt.Println("t is now", t)
上面代码的输出如下：
t is now {77 Sunset Strip}
如果变量 s 是通过 t ，而不是 &t 创建的，调用 SetInt 和 SetString 将会失败，因为 t 的字段不是“可写的”。