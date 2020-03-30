## defer的使用

> golang语言中defer的使用场景较多，用于锁的关闭，连接的延迟关闭等，通常在函数的结束时调用，简单的说就是`在函数结束时 ` `返回值赋值后`，`返回前执行defer的方法,最后才返回`。

> 另外defer在go 1.13版本（包括）前确实有一定的开销，拒绝滥用，但是go 1.14版本（包括）后，官方给出的回应是：Go1.14提高了defer的大多数用法的性能，几乎0开销！defer已经可以用于对性能要求很高的场景了。

> 通过defer，我们可以在代码中优雅的关闭/清理代码中所使用的变量。defer作为golang清理变量的特性，有其独有且明确的行为。以下是defer三条使用规则。


### 规则一 当defer被声明时，其参数就会被实时解析
 
举例：
 
 ```go
 func DeferTestA() {
 	i := 0
 defer fmt.Printf("defer 输出的 i 值：%d \n ",i)
 	i++
 	return
 }
```
运行结果是 `defer 输出的 i 值：0 `

这是因为虽然我们在defer后面定义的是一个带变量的函数: `fmt.Printf("defer 输出的 i 值：%d \n ",i)` 。但这个变量`i`在defer被声明的时候，就已经确定其值了。 上面的代码等同于下面的代码:

```go
func DeferTestA() {
	i := 0
	defer fmt.Printf("defer 输出的 i 值：%d \n ",0)
	i++
	return
}
```

为了更明确的说明这个问题，我们再定义一个defer:

```go
func DeferTestB() {
	i := 0
	defer fmt.Printf("defer 第一次输出的 i 值：%d \n ",i)
	i++
	defer fmt.Printf("defer 第二次输出的 i 值：%d \n ",i)
	return
}
```

运行结果：
```go 
defer 第二次输出的 i 值：1 
defer 第一次输出的 i 值：0 
```

通过运行结果，可以看到defer输出的值，就是定义时的值。
并且，如果你注意到结果顺序，你会发现输出顺序和defer顺序相反，这就是接下来要说的规则二。

### 规则二 defer执行顺序为先进后出

上面的例子已经可以说明这一问题。`defer 不带函数执行时`,defer可以理解像栈，先进后出。
当同时定义了多个defer代码块时，golang 先定义后执行的顺序依次调用defer。不要问为什么，golang就是这么定义的。为了加深记忆和理解，用下面的例子来加以说明:

```go

func DeferTestC() {
	for i := 0; i < 4; i++ {
		defer fmt.Printf("defer 第 %d 次输出的 i 值：%d \n ",i,i)
	}
}
```

运行结果：
```go
defer 第 3 次输出的 i 值：3 
 defer 第 2 次输出的 i 值：2 
 defer 第 1 次输出的 i 值：1 
 defer 第 0 次输出的 i 值：0 

```

在循环中，相当于依次定义了四个defer代码块。结合规则一，我们可以明确得知每个defer代码块应该输出什么值。 根据先进后出的原则，我们可以看到依次输出了 3 2 1 0 。


### 规则三 defer可以读取有名返回值，defer函数执行在`return` 调用后返回值前

前面的例子都是不带函数的，当 `defer带函数执行时`，defer函数执行是在return返回赋值后，返回前执行，看下面的例子：

下面的例子，我们定义 一个返回值`res = 1`，然后返回`return` res时，赋值为 7 ，根据刚才所讲，defer函数执行在返回赋值后，返回前，那么此时res为7，defer 函数执行将res赋值为8，return最后返回应该为8
看下面代码：

package main
```go
func DeferTestD() (res int) {
	res = 1
	defer fmt.Println("one")  //第一个进入
	defer fmt.Println("two")  //第二进入
	defer fmt.Println("three") //第三个进入
	defer func() {
		fmt.Println("start", res)
		res++
		fmt.Println("end", res)
	}()
	return 7
}
```

运行结果：
```go
start 7
end 8
three
two
one
8
```

根据上面的例子来看,`defer`一定是延迟执行的么？答案是肯定的。但是有注意点需要注意：

如下面的例子：

```go
type SliceNum []int

func NewSlice() SliceNum {
	return make(SliceNum, 0)

}

func (s *SliceNum) Add(elem int) *SliceNum {
	*s = append(*s, elem)
	fmt.Println("add", elem)
	fmt.Println("add SliceNum end", s)
	return s
}

func main(){
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(10)
	s.Add(11)
	s.Add(12)
}
```
运行结果：
```go
add 1
add SliceNum end &[1]
add 10
add SliceNum end &[1 10]
add 11
add SliceNum end &[1 10 11]
add 12
add SliceNum end &[1 10 11 12]
add 2
add SliceNum end &[1 10 11 12 2]
```

此时defer后面的一个方法Add(1)先执行了，然后才执行的s.Add(10),s.Add(11),s.Add(12),最后才执行了Add(2)

我们再改变下main函数如下：

```go
func main(){
	s := NewSlice()
	defer s.Add(1).Add(2).Add(3).Add(4)
	s.Add(10)
	s.Add(11)
	s.Add(12)
}
```
运行结果如下：
```go
add 1
add SliceNum end &[1]
add 2
add SliceNum end &[1 2]
add 3
add SliceNum end &[1 2 3]
add 10
add SliceNum end &[1 2 3 10]
add 11
add SliceNum end &[1 2 3 10 11]
add 12
add SliceNum end &[1 2 3 10 11 12]
add 4
add SliceNum end &[1 2 3 10 11 12 4]
```

可以看到`defer`中先执行的s.Add(1).Add(2).Add(3)，然后执行s.Add(10)， s.Add(11)，s.Add(12)，延迟执行的test函数，可以看到defer延迟执行的是最后的一个函数add(4)

如何保证整个defer是在最后执行呢？当然可以了使用 defer fun(){
}()，包住此方法。
举例：
```go
func main(){
	s := NewSlice()
	defer func() {
		s.Add(1).Add(2).Add(3).Add(4)
	}()
	s.Add(10)
	s.Add(11)
	s.Add(12)
}
```
运行结果：
```go
add 10
add SliceNum end &[10]
add 11
add SliceNum end &[10 11]
add 12
add SliceNum end &[10 11 12]
add 1
add SliceNum end &[10 11 12 1]
add 2
add SliceNum end &[10 11 12 1 2]
add 3
add SliceNum end &[10 11 12 1 2 3]
add 4
add SliceNum end &[10 11 12 1 2 3 4]
```
