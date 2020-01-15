## 字符串(string)长度

### 计算字符串长度

> Go语言的内建函数len()，可以用来获取切片、字符串 、通道(channel)等的长度。注：返回的数据类型是 ```int``` 型。


下面的代码可以用len()来获取字符串的长度。 

```go
    str_1 :="hello go!"
    fmt.Println(len(str_1))
    str_2 :="你好，go！"
    fmt.Println(len(str_2))
```
程序输出如下 :
```$xslt
9
14
```

len()函数的返回值的类型为int,表示字符串的ASCII字符个数或字节长度。
- 输出中第一行的9表示str_1的字符个数为9
- 输出中第二行的14表示str_2的字符格式，也就是"你好，go！"的字符个数是14，然而根据习惯， "你好，go！"的字符个数应该是6

这里的差异是由于Go语言的字符串都以UTF-8格式保存，每个中文占用3个字节，因此使用len()获得两个中文文字和两个中文字符（，与！）对应的12个字节。

如果希望按习惯上的字符个数来计算，就需要使用Go语言中UTF-8包提供的```RuneCountInString()```函数来统计Uncode字符数量。

下面的代码展示如何计算UTF-8的字符个数。
```go
//计算含中文的字符串长度
fmt.Println("计算含中文的字符串长度: ",utf8.RuneCountInString("你好，go！"))
```
程序输出如下 :
```$xslt
计算含中文的字符串长度:  6
```

总结 :
- ASCII字符串长度使用len()函数。
- Unicode字符串长度使用utf8.RuneCountInString()函数。

