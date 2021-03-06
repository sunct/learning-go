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

type Student struct {
	Name string
	Age int
}


var p1 = Student{Name:"WangXiaoMing",Age:18}

func FmtDemo(){
	fmt.Printf("\n %v", p1)
	fmt.Printf("\n %+v", p1)
	fmt.Printf("\n %#v", p1)
	fmt.Printf("\n %T", p1)
	fmt.Printf("\n %%")
	fmt.Printf("\n %t",!true)

	fmt.Printf("\n %b",100)
	fmt.Printf("\n %c",0x4E2D)
	fmt.Printf("\n %c",0x56FD)
	fmt.Printf("\n %d",0100)
	fmt.Printf("\n %o",64)
	fmt.Printf("\n %q",65)
	fmt.Printf("\n %q",0x56FD)
	fmt.Printf("\n %x",123)
	fmt.Printf("\n %X",123)
	fmt.Printf("\n %U",0x56FD)

	fmt.Printf("\n %s", []byte("中国加油！"))
	fmt.Printf("\n %q", "中国加油！！")
	fmt.Printf("\n %x", "golang")
	fmt.Printf("\n %X", "golang")
	fmt.Printf("\n %p", &p1)
}


func DeferTestA() {
	i := 0
	defer fmt.Printf("defer 输出的 i 值：%d \n ",i)
	i++
	return
}

func DeferTestB() {
	i := 0
	defer fmt.Printf("defer 第一次输出的 i 值：%d \n ",i)
	i++
	defer fmt.Printf("defer 第二次输出的 i 值：%d \n ",i)
	return
}

func DeferTestC() {
	for i := 0; i < 4; i++ {
		defer fmt.Printf("defer 第 %d 次输出的 i 值：%d \n ",i,i)
	}
}
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

	//TestStringLen()
	//TestTypeConversion()
	//FmtDemo()
	//fmt.Println(DeferTestD())
	//fmt.Println(test())
	//s := NewSlice()
	//defer func() {
	//	s.Add(1).Add(2).Add(3).Add(4)
	//}()
	//s.Add(10)
	//s.Add(11)
	//s.Add(12)

	fmt.Printf("%g",fmt.Sprintf("%f",99999999.99))
	//fmt.Sprintf("%g", fmt.Sprintf("%f", 99999999.99))
}


