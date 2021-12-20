/*
 * @Author: vyron
 * @Date: 2021-06-18 13:06:29
 * @LastEditTime: 2021-08-19 19:21:15
 * @LastEditors: vyron
 * @Description: go string
 * @FilePath: /Golang/record/02.basic-grammar/04.string/main.go
 */

package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

/**

字符串是一种值类型，且值不可变，即创建某个文本后将无法再次修改这个文本的内容，更深入地讲，字符串是字节的定长数组。

定义字符串

使用双引号定义字符串,可使用转义字符实现换行,缩进
\n: 换行符 new line
\r: 回车符 return
\t: 缩进 tab
\u或\U: Unicode字符
\\: 反斜杠

检测字符串长度时,ASCII 字符串长度使用 len() 函数, Unicode 字符串长度使用 utf8.RuneCountInString() 函数。

遍历字符串时,ASCII字符串使用下标遍历,Unicode字符串使用 for range 遍历

Go 语言的字符串是不可变的。修改字符串时，可以将字符串转换为 []byte 进行修改。[]byte 和 string 可以通过强制类型转换互转。


*/

func main(){
	// defineString();
	// spliceString();
	// getStringLen();
	// traversalString();
	// splitString();
	// changeString();
	// connectString();
	formatString();
}

// 定义字符串
func defineString(){
	var str = "Learn go\nyou will be happy!";
	length:=len(str);
	fmt.Println(str);
	fmt.Println("字符串长度为:",length);
	// for i := 0; i < length; i++ {
	// 	fmt.Println(i,":",str[i]);
	// }

	// 定义多行字符串,转义字符无效
	var str2 = `第一行
第二行
第三行
\n
\n
\n
...`;
	fmt.Println(str2);
}


// 拼接字符串
func spliceString(){
	var str = "hello "+"world";
	str += " golang";
	fmt.Println(str);
}

// 获取字符串长度
func getStringLen(){
	var str = "Learn golang, you will be happy, 哈哈哈,我非常开心!";
	// go内置函数len(),用来获取切片、字符串、通道（channel）等的长度,表示字符串的ASCII字符个数或字节长度
	fmt.Println("ASCII字节长度:",len(str));
	// go字符串以 UTF-8 格式保存,中文占用三个字节
	fmt.Println("字符串长度:",utf8.RuneCountInString(str));
}

// 字符串遍历
func traversalString(){
	var str = "Learn golang, you will be happy, 哈哈哈,我非常开心!";
	// ASCII遍历
	for i := 0; i < len(str); i++ {
		fmt.Printf("ASCII: %c  %d\n",str[i],str[i]);
	}
	// Unicode字符遍历
	for _, s := range str {
		fmt.Printf("Unicode: %c  %d\n",s,s);
	}
}

// 截取字符串
func splitString(){
	str := "Hello World!";
	// 正向搜索字符串  strings.LastIndex() 反向搜索字符串
	index := strings.Index(str," ");
	// str[:index] 返回str字符串index位置前的字符串,str[index:]返回index位置后的字符串
	fmt.Println("index:",index,str[:index]);
}

// 修改字符串
func changeString(){
	// 英文
	str:="abc" 
	s2:=[]byte(str) 
	s2[0]='b' 
	fmt.Println(string(s2))

	// 中文
	str:="白猫" 
	s2:=[]rune(str) 
	s2[0]='黑' 
	fmt.Println(string(s2))
}

// 字符串拼接
func connectString(){
	hello := "hello";
	world := " world";
	// 使用 + 拼接字符串
	fmt.Println(hello+world);
	// 使用bytes.Buffer拼接字符串(高效)
	var stringButter bytes.Buffer;
	stringButter.WriteString(hello);
	stringButter.WriteString(world);
	fmt.Println(stringButter.String());
}

// 格式化字符串
func formatString(){
	finish:=100;
	current:=20;
	progress:=fmt.Sprintf("加载中,当前进度:%d,总进度:%d",current,finish);
	fmt.Println(progress);


	pi:=3.1415926;
	convertPi := fmt.Sprintf("%v %v %v","圆周率",pi,true);
	fmt.Println(convertPi);

	// 其他格式样式可在fmt库中详细查看
}