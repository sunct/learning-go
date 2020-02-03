## fmt格式化输出（附占位符）待补充

> 在开发调试过程中，避免不了输出打印一些信息，以便查看相关数据的格式和正确性。

常用的函数```fmt.Printf()```,```fmt.Println```



>附 占位符说明

以下结构体为例
```go
type Student struct {
	Name string
	Age int
}

var p1 = Student{Name:"WangXiaoMing",Age:18}
```
- 普通占位符

| 占位符 | 说明 |举例|
| :----:|:--- |:-----|
|%v     | 相应值的默认格式。 |`fmt.Printf("\n %v", p1)` <br/>结果：{WangXiaoMing 18} |
|%+v    | 在打印结构体时，“加号”标记（%+v）会添加字段名|`fmt.Printf("\n %+v", p1)` <br/>结果：{Name:WangXiaoMing Age:18}|
|%#v    | 相应值的Go语法表示 |`fmt.Printf("\n %#v", p1)`<br/>结果：main.Student{Name:"WangXiaoMing", Age:18} |
|%T     | 相应值的类型的Go语法表示|`fmt.Printf("\n %T", p1)`<br/>结果：main.Student|
|%%     | 字面上的百分号，并非值的占位符|`fmt.Printf("\n %%")`<br/>结果：%|


- 布尔占位符

| 占位符 | 说明 |举例|
| :----:|:--- |:-----|
|%t     | true 或 false|`fmt.Printf("\n %t",!true)` <br/>结果：false|


- 整数占位符

| 占位符 | 说明 |举例|
| :----:|:--- |:-----|
|%b     | 二进制表示  | `fmt.Printf("\n %b",100)`<br/>结果： 1100100|
|%c     | 相应Unicode码点所表示的字符 | `fmt.Printf("\n %c",0x56FD)`<br/>结果：国|
|%d     | 十进制表示  |`fmt.Printf("\n %d",0100)`<br/>结果：64 |
|%o     | 八进制表示  |`fmt.Printf("\n %o",64)` <br/> 结果：100 |
|%q     | 单引号围绕的字符字面值，由Go语法安全地转义 | `fmt.Printf("\n %q",65)` <br/>结果：'A'|
|%x     | 十六进制表示，字母形式为小写 a-f |`fmt.Printf("\n %x",123)`<br/>结果：7b |
|%X     | 十六进制表示，字母形式为大写 A-F |`fmt.Printf("\n %X",123)`<br/>结果：7B |
|%U     | Unicode格式：U+1234，等同于 "U+%04X"|`fmt.Printf("\n %U",0x56FD)`<br/>结果：U+56FD |   


- 浮点数和复数的组成部分（实部和虚部）

| 占位符 | 说明 |举例|
| :----:|:--- |:-----|
|%b     | 无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat的 'b' 转换格式一致。|例如 -123456p-78|
|%e     | 科学计数法，例如 -1234.456e+78 |                           Printf("%e", 10.2)   |
|%E     | 科学计数法，例如 -1234.456E+78  |                         Printf("%e", 10.2)     |
|%f     | 有小数点而无指数，例如 123.456   |                        Printf("%f", 10.2)      |
|%g     | 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出 |       Printf("%g", 10.20)   |
|%G     | 根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出  |      Printf("%G", 10.20+2i) |

- 字符串与字节切片

| 占位符 | 说明 |举例|
| :----:|:--- |:-----|
|%s     | 输出字符串表示（string类型或[]byte) |    `fmt.Printf("\n %s", []byte("中国加油！"))`<br/> 结果：中国加油！ |
|%q     | 双引号围绕的字符串，由Go语法安全地转义 |   `fmt.Printf("\n %q", "中国加油！！")`<br/>结果："中国加油！！")   | 
|%x     | 十六进制，小写字母，每字节两个字符   |     `fmt.Printf("\n %x", "golang")`<br/>结果：676f6c616e67)      |   
|%X     | 十六进制，大写字母，每字节两个字符   |     `fmt.Printf("\n %X", "golang")`<br/> 结果：676F6C616E67|

- 指针

| 占位符 | 说明 |举例|
| :----:|:--- |:-----|
|%p     | 十六进制表示，前缀 0x |  `fmt.Printf("\n %p", &p1)`<br/>结果：0x1174c10|



参考来源：[官网](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.3.html)