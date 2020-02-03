## 类型(int,int64,string等)之间相互转换

> Go语言是强类型的编程语言，在编码过程中避免不了有些类型的转换操作，如果前后类型不一致就会报错。本篇主要汇总一些转换的方法和函数，以下所说的字符串，一般都是由数字组成的字符串，如"202001015"。

>  附：各int类型的取值范围如下:
```$xslt
 int8: -128 ~ 127
 int16: -32768 ~ 32767
 int32: -2147483648 ~ 2147483647
 int64: -9223372036854775808 ~ 9223372036854775807
 uint8: 0 ~ 255
 uint16: 0 ~ 65535
 uint32: 0 ~ 4294967295
 uint64: 0 ~ 18446744073709551615
```
> 注：使用自动推到 `:=` 的整型数据都是 `int`型。
- ### int 转 字符串
使用函数 `strconv.Itoa()`

例如：
```go
	// int到string
	int_v:=11223344
	str_0 := strconv.Itoa(int_v)
	fmt.Printf("\n int转换后的类型 %T ",str_0)
	fmt.Printf("\n int转换后的数值 %s ",str_0)

```
打印输出值为：
```go
 int转换后的类型 string 
 int转换后的数值 11223344 
```

- ### 字符串转 int
使用函数 `strconv.Atoi()`
```go
	// string到int
	str_1:="12344321"
	need_int, err_0 := strconv.Atoi(str_1)
	if err_0!=nil{
		fmt.Println("转换出错",err_0)
	}
	fmt.Printf("\n 字符串转换后的类型 %T ",need_int)
	fmt.Printf("\n 字符串转换后的数值 %d ",need_int)
```
打印输出值为：
```go
 字符串转换后的类型 int 
 字符串转换后的数值 12344321
```
- ### int64 转 字符串

使用函数 `strconv.FormatInt()`

例如：
```go
    //int64到string
    var int64_v int64  = 123454321
    string := strconv.FormatInt(int64_v,10)
    fmt.Printf("\n int64转换后的类型 %T ",string)
    fmt.Printf("\n int64转换后的数值 %s ",string)
```
打印输出值为：
```go
 int64转换后的类型 string 
 int64转换后的数值 123454321 
```

- ### 字符串 转 int64
使用函数 `strconv.ParseInt()`

例如：
```go
    //string 转int64
    str:="1234567890"
    need_int64, err := strconv.ParseInt(str, 10, 64)
    if err!=nil{
    	fmt.Println("转换出错",err)
    }
    fmt.Printf("\n 字符串转换后的类型 %T ",need_int64)
    fmt.Printf("\n 字符串转换后的数值 %d ",need_int64)
```
打印输出值为：
```go
 字符串转换后的类型 int64 
 字符串转换后的数值 1234567890 
```
- ### int 转 int64 
 同是整型，可以通过强制转换,使用 `int64()`
 ```go
	//int 转 int64
	int_v2:=12345
	int64_v2:=int64(int_v2)
	fmt.Printf("\n int转换后的类型 %T ",int64_v2)
	fmt.Printf("\n int转换后的数值 %d ",int64_v2)
```
打印输出值为：
```go
 int转换后的类型 int64 
 int转换后的数值 12345 
```

- ### int64 转 int
 同是整型，可以通过强制转换,使用 `int()`
```go
	//int64 转 int
	var int_v3 int64 =1234567
	int64_v3:=int(int_v3)
	fmt.Printf("\n int64转换后的类型 %T ",int64_v3)
	fmt.Printf("\n int64转换后的数值 %d ",int64_v3)
```
打印输出值为：
```go
 int64转换后的类型 int 
 int64转换后的数值 1234567
```

其他类型间转换待随后使用时补充。