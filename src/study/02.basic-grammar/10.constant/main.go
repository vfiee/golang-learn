package main

import (
	"fmt"
	"math"
)

/*
Go语言常量

Go 语言常量使用 const 关键字定义,用于存储不会修改的数据
1. 常量编译时创建,既是在函数内部
2. 只能是布尔型,数字型(整数,浮点数,复数),字符串型
3. 常量表达式必须是能被编译器求值的常量表达式

*/


func main(){
	constIota();
	noTypeConst();
}

type Weekday int;

func constIota(){
	// iota 常量生成器
	const (
		Sunday Weekday = iota
		Monday
    	Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(
		Sunday,
		Monday,
    	Tuesday,
		Wednesday,
		Thursday,
		Friday,
		Saturday,
		);
}


// 编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算，可以认为至少有 256bit 的运算精度。
// 这里有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。
func noTypeConst(){
	// math.Pi 无类型的浮点数常量，可以直接用于任意需要浮点数或复数的地方
	var x float32 = math.Pi;
	var y float64 = math.Pi;
	var z complex128 = math.Pi;

	fmt.Println(x,y,z)
}