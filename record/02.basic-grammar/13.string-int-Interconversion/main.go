package main

import (
	"fmt"
	"strconv"
)

/*
Go字符串和数字相互转换


*/


func main(){
	stringToInt("1000")
	intToString(100)
	boolSeries()
	intSeries()
}
func stringToInt(s string){
	i,err := strconv.Atoi(s);
	if err != nil {
		fmt.Println("stringToInt转换失败");
	}else {
		fmt.Println("stringToInt转换成功",i);
	}
}

func intToString(i int){
	s := strconv.Itoa(i);
	fmt.Println("intToString转换成功",s);
}

func boolSeries(){
	fmt.Println("bool series")
	appendBool();
	formatBool();
	parseBool();
}

func appendBool(){
	b := []byte("bool:")
	bt := strconv.AppendBool(b, true)
	bf := strconv.AppendBool(b, false)
	fmt.Println(string(bt),string(bf))
}
func formatBool(){
	// 将bool类型转为字符串类型
	fmt.Println("formatBool(false):",strconv.FormatBool(false));
	fmt.Println("formatBool(true):",strconv.FormatBool(true));
}
func parseBool(){
	t := "true";
	f := "false";
	tr,terr := strconv.ParseBool(t);
	fr,ferr := strconv.ParseBool(f);
	if terr!=nil {
		fmt.Println(tr);
	}
	if(ferr != nil) {
		fmt.Println(fr);
	}
}

func intSeries(){
	fmt.Println("int series")
	appendInt()
	formatInt()
	parseInt()
}

func appendInt() {
	b10 := []byte("int (10进制):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println("appendInt",string(b10))

	b16 := []byte("int (16进制):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println("appendInt",string(b16))
}
func formatInt() {
	v := int64(-42)
	s10 := strconv.FormatInt(v, 10)
	s16 := strconv.FormatInt(v, 16)
	fmt.Println("formatInt十进制:",s10);
	fmt.Println("formatInt十六进制:",s16);
}
func parseInt() {
	v32 := "-354634382"
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Println("parseInt十进制:32",s)
	}
	if s, err := strconv.ParseInt(v32, 16, 64); err == nil {
		fmt.Println("parseInt十六进制:32",s)
	}else {
		fmt.Println(err);
	}

	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Println("parseInt十进制:64",s)
	}
	if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
		fmt.Println("parseInt十六进制:64",s)
	}else {
		fmt.Println(err);
	}
}