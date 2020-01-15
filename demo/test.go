/**
 * @Author: sunct
 * @Description: 
 * @File:  test.go
 * @Version: 1.0.0
 * @Date: 2020-01-15 11:08
 */
package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

/*
 * 测试字符串长度
 */
func TestStringLen(){

	str_1 :="hello go!"
	fmt.Println(len(str_1))
	str_2 :="你好，go！"
	fmt.Println(len(str_2))

	//计算含中文的字符串长度
	fmt.Println("计算含中文的字符串长度: ",utf8.RuneCountInString("你好，go！"))

}
/*
 * 测试类型互转
 */
func  TestTypeConversion(){

	// int到string
	int_v:=11223344
	str_0 := strconv.Itoa(int_v)
	fmt.Printf("\n int转换后的类型 %T ",str_0)
	fmt.Printf("\n int转换后的数值 %s ",str_0)

	// string到int
	str_1:="12344321"
	need_int, err_0 := strconv.Atoi(str_1)
	if err_0!=nil{
		fmt.Println("转换出错",err_0)
	}
	fmt.Printf("\n 字符串转换后的类型 %T ",need_int)
	fmt.Printf("\n 字符串转换后的数值 %d ",need_int)

	//string 转int64
	str:="1234567890"
	need_int64, err := strconv.ParseInt(str, 10, 64)
	if err!=nil{
		fmt.Println("转换出错",err)
	}
	fmt.Printf("\n 字符串转换后的类型 %T ",need_int64)
	fmt.Printf("\n 字符串转换后的数值 %d ",need_int64)

	//int64到string
	var int64_v int64  = 123454321
	string := strconv.FormatInt(int64_v,10)
	fmt.Printf("\n int64转换后的类型 %T ",string)
	fmt.Printf("\n int64转换后的数值 %s ",string)

	//int 转 int64
	int_v2:=12345
	int64_v2:=int64(int_v2)
	fmt.Printf("\n int转换后的类型 %T ",int64_v2)
	fmt.Printf("\n int转换后的数值 %d ",int64_v2)
	//int64 转 int
	var int_v3 int64 =1234567
	int64_v3:=int(int_v3)
	fmt.Printf("\n int64转换后的类型 %T ",int64_v3)
	fmt.Printf("\n int64转换后的数值 %d ",int64_v3)


}

func main(){

	//TestStringLen()
	TestTypeConversion()

}

