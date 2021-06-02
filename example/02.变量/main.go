package main

// 所有内存在Go中都是经过初始化的
// 当变量被声明后,系统自动赋值变量的零值: int为0, float为0.0, bool为false, string为空字符串,
// 切片、函数、指针变量的默认为 nil

/*
基本数据类型

bool
string
int int8 int16 int32 int64
unit unit8 unit16 unit32 unitptr
byte: unit8 的别名
rune: int32的别名 代表一个 Unicode 码
float32 float64
complex64 complex128

*/

/*

变量的声明

单个声明变量:
var 变量名 变量类型


批量声明变量:
var (
	testInt int,
	testString string,
	testFloat []float32
	testFunc func() bool
	testStruct struct {
		testBool bool
	}
)

简短声明变量:
func main(){
	testString:="I\'m string"
	testInt,testBool:="string",true
}

简短声明变量的限制
1. 只能在函数内部
2. 定义变量的同时,显式初始化变量
3. 不能提供数据类型

*/

/*

变量的初始化

标准初始化:
var 变量名 类型 = 表达式

编译器推导类型初始化: 根据 = 右侧表达式的值推导变量类型
var 变量名 = 表达式
var 变量1, 变量2 = 表达式1, 表达式2

短变量声明并初始化: 编译器根据 = 右侧值得类型推断出左侧对应类型
变量名 := 表达式
变量1, 变量2 := 表达式1, 表达式2

短变量的限制:
1. 短变量声明单个变量时, 变量名必须是新变量,否则编译报错: no new variables on left side of :=
2. 短变量声明多个变量时, 至少其中一个变量为新变量,即使其他变量名可能重复声明,编译器也不会报错

*/

/*

变量的赋值

变量交换

int1 := 100
int2 := 200

int2, int1 = int1, int2

*/

/*

匿名变量
1. 下划线 "_" 被称为空白标识符,使用下划线代表匿名变量
2. 任何类型的值都可以赋值给匿名变量, 但匿名变量的任何值都将被抛弃
3. 匿名变量不能作为变量与其他变量进行赋值或运算
4. 匿名变量不分配内存,不占用内存

func getData(){
	return "string", 100
}

func main(){
	str, _ := getData()
	_, number := getData()
	fmt.Println(str,number)
}

*/

/*

变量的作用域:
1. 函数内定义的变量称为局部变量(函数参数和返回值变量都是局部变量)
func main(){
	a := 3;
	b := 4;
	c := a + b;
	fmt.Printf("a = %d, b = %d, c = %d \n",a,b,c)
}
2. 函数外定义的变量称为全局变量(函数体内,局部变量可以同全局变量相同,函数体内的局部变量优先被考虑)
var c int;

func main(){
	a := 3;
	b := 4;
	c = a + b;
	fmt.Printf("a = %d, b = %d, c = %d \n",a,b,c)
}
3. 函数定义中的变量称为形式参数
函数定义时接受的变量称为形式参数,形参在函数调用时生效,调用结束后销毁,未调用时不占用实际存储单元;

func main(){
	a := 3;
	b := 4;
	c := sum(a,b);
	fmt.Printf("a = %d, b = %d, c = %d \n",a,b,c)
}

func sum(a,b int) int{
	return a + b;
}

*/

