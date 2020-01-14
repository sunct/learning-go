## golong -【包的使用】汉语转拼音


> 业务使用背景
给定汉字自动生成拼音首字母，按纯汉字生成拼音小写首字母（包含的英文、数字等忽略），如：王小&and二放牛qule -> wxefn
!!!

### 1、包的选用

我一开始引用的包是：
``` git
https://github.com/sunct/go-pinyin
```
该包的功能满足当前业务应用，但是，运维人员在部署[[Docker]]的时候出现以下问题：

error Error

>github.com/mozillazg/go-pinyin
/data/golang/go_repos/pkg/mod/github.com/mozillazg/go-pinyin@v0.16.0/phrase.go:10:10: undefined: gojieba.NewJieba


经过排查，发现包中含有c语言,打包失败。
为了能更快的部署，于是换了一个包：
```git
https://github.com/go-ego/gpy
```
> 该包已fork到我的git([sunct (sunct) · GitHub](https://github.com/sunct))，也可使用：

```git
https://github.com/sunct/gpy
```
可用模式：
```go
Normal      = 0 // 普通风格，不带声调（默认风格）。如： zhong guo
Tone        = 1 // 声调风格1，拼音声调在韵母第一个字母上。如： zhōng guó
Tone2       = 2 // 声调风格2，即拼音声调在各个韵母之后，用数字 [1-4] 进行表示。如： zho1ng guo2
Tone3       = 8 // 声调风格3，即拼音声调在各个拼音之后，用数字 [1-4] 进行表示。如： zhong1 guo2
Initials    = 3 // 声母风格，只返回各个拼音的声母部分。如： zh g
FirstLetter = 4 // 首字母风格，只返回拼音的首字母部分。如： z g
Finals      = 5 // 韵母风格，只返回各个拼音的韵母部分，不带声调。如： ong uo
FinalsTone  = 6 // 韵母风格1，带声调，声调在韵母第一个字母上。如： ōng uó
FinalsTone2 = 7 // 韵母风格2，带声调，声调在各个韵母之后，用数字 [1-4] 进行表示。如： o1ng uo2
FinalsTone3 = 9 // 韵母风格3，带声调，声调在各个拼音之后，用数字 [1-4] 进行表示。如： ong1 uo2

```

### 2、代码处理

引用包中已经返回拼音部分,封装函数：
```go
/**
 * @Author: sunct
 * @Description:
 * @File:  pinyinCode.go
 * @Version: 1.0.0
 * @Date: 2020-01-10 13:06
 */
package pinyinCode

import (
	"bytes"
	"github.com/go-ego/gpy"
	"reflect"
)

type PinyinCode struct {
}

var ChPinyinCode = newPinyinCode()

func newPinyinCode() *PinyinCode {
	return &PinyinCode{}
}

/*直接转 汉字转拼音（已忽略非汉字）
 * sunct
 */
func GetPinyinCode(zh string) (pyCode string) {
	var buffer bytes.Buffer
	a := gpy.NewArgs()
	//首字母模式
	a.Style = gpy.FirstLetter
	returnData := gpy.Pinyin(zh, a)
	////处理返回的数据，取出首字母拼接成字符串
	//// 如 zh = "汉字拼音"   返回[[h] [z] [p] [y]] 拼接成 hzpy
	for _, v := range returnData {
		pinyin := v[0]
		buffer.WriteString(pinyin)
	}
	return buffer.String()
}

/* 根据结构体转 汉字转拼音
 * 参数 1 结构体数据 2 需要根据转换的字段 3 转换到的字段
 * sunct
 * 举例（只是个例子）：
	//map
    need := map[string]string{
		"PinyinCode": "Name", // PinyinCode 字段使用 Name 转成的拼音
		"Unit":       "Remark",// Unit 字段使用 Remark 转成的拼音
	}

	//调用
	err = pinyinCode.ChPinyinCode.ChangePinyinCode(in, need)
	... ...
*/
//
func (this *PinyinCode) ChangePinyinCode(in interface{}, needle map[string]string) (err error) {
	var t reflect.Type
	t = reflect.TypeOf(in)
	v := reflect.ValueOf(in)
	// 普通类型不解析，直接返回
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Ptr && v.Kind() != reflect.Struct {
		return nil
	}
	_, err = this.parseParam(t, v, needle)
	return err

}
func (this *PinyinCode) parseParam(t reflect.Type, v reflect.Value, needle map[string]string) (out interface{}, err error) {
	// 判断是否是数组
	if v.Kind() == reflect.Slice {
		//fmt.Print("是数组")
		for i, n := 0, v.Len(); i < n; i++ {
			t1 := reflect.TypeOf(v.Index(i).Interface())
			_, err := this.parseParam(t1, v.Index(i), needle)
			if err != nil {
				return nil, err
			}
		}
		// 是否是指针
	} else if v.Kind() == reflect.Ptr {
		//fmt.Print("是指针")
		t = t.Elem()
		v = v.Elem()
		data, err := this.parseParam(t, v, needle)
		if err != nil {
			return nil, err
		}
		return data, nil
		// 结构体的
	} else if v.Kind() == reflect.Struct {
		//fmt.Print("是结构体")
		for i := 0; i < t.NumField(); i++ {
			//获取 needle 的键值对信息
			for k0, v0 := range needle {
				//获取初始汉字字段
				zh := v.FieldByName(v0).String()

				if v.FieldByName(v0).IsValid() {
					//把汉字转成拼音，并赋值
					pinyinCode := GetPinyinCode(zh)
					v.FieldByName(k0).Set(reflect.ValueOf(pinyinCode))
				}
			}
		}
		return v, nil
		// 最后普通类型解析
	} else {
		panic("兄弟er，你走错包间了")
	}
	return nil, nil
}




```


### 3、方法使用

[[GetPinyinCode]] 函数可直接调用
例如；
```go
py:= pinyinCode.GetPinyinCode("这是汉字字符串")
```


[[ChangePinyinCode]] 函数使用前先定义一个[[map]]

例如：
```go
need := map[string]string{
        "PinyinCode": "Name", // PinyinCode 字段使用 Name 转成的拼音
        "Unit":       "Remark",// Unit 字段使用 Remark 转成的拼音
    }

    //调用
    err = pinyinCode.ChPinyinCode.ChangePinyinCode(in, need)
    ... ...
  ```
  其中，in 参数是需要的结构体数据（注：里面必须包括map中用到的所有字段名称），通过本函数处理返回后的数据中已把相关字段赋值给相关字段（通过指针处理了），可直接使用数据（比如：in）。
  




