/**
 * @Author: sunct
 * @Description: 
 * @File:  test.go
 * @Version: 1.0.0
 * @Date: 2020-01-15 11:08
 */
package demo

import (
	"fmt"
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

func main(){

	TestStringLen()

}

