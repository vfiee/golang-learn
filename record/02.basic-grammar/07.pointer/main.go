package main

import (
	"flag"
	"fmt"
)

/*
Go 语言指针

指针被拆分为两个核心概念:

1. 类型指针: 允许对这个指针类型的数据进行修改,传递数据可直接使用指针,无需拷贝数据,类型指针不能进行偏移和运算
2. 切片: 由指向起始元素的原始指针,元素数量和容量组成

切片比原始指针具备更强大的特性，而且更为安全。切片在发生越界时，运行时会报出宕机，并打出堆栈，而原始指针只会崩溃。



指针地址,指针类型,指针取值

指针地址:
1. 一个指针变量可以指向任意一个值得内存地址;
2. 指针变量指向得内存地址在 32 和 64 位机器上分别占用 4 或 8 字节,占用字节大小与所指向的值大小无关;
3. 指针被定义后未分配任何变量时,默认值未 nil
4. 指针变量通常缩写未 ptr
5. 每个变量运行时都有一个地址表示变量在内存中的位置,在变量前添加 & 操作符可以获取变量内存地址;


n := 32; 声明类型为 int 的变量 n,并赋值 32
ptr := &n; 声明类型为 *int 的变量, 并赋值
value := *ptr; 声明类型为 指针类型值 的变量,并赋值

ptr 的类型为 *int,称为 int 的指针类型, * 代表指针;

取值操作符 * 和取地址操作符是一对互补操作符,&取出地址,*取出地址指向的值;



变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
1. 对变量进行取地址操作使用&操作符，可以获得这个变量的指针变量。
2. 指针变量的值是指针地址。
3. 对指针变量进行取值操作使用*操作符，可以获得指针变量指向的原变量的值。



*/


func main(){
	// getPointerAddress();
	// getPointerValue();
	// swap();
	// flagParseCliArgs();
	newPointerType();
}

func getPointerAddress(){
	n := 32;
	ptr := &n;
	fmt.Println("ptr:",ptr);
}

func getPointerValue(){
	// 准备一个字符串类型
    var house = "Malibu Point 10880, 90265"
    // 对字符串取地址, ptr类型为*string
    ptr := &house
    // 打印ptr的类型
    fmt.Printf("ptr type: %T\n", ptr)
    // 打印ptr的指针地址
    fmt.Printf("ptr address: %p\n", ptr)
    // 对指针进行取值操作
    value := *ptr
    // 取值后的类型
    fmt.Printf("value type: %T\n", value)
    // 指针取值后就是指向变量的值
    fmt.Printf("value: %s\n", value)
}

func swap(){
	x,y := 1,2;

	// swap
	t := *&x;
	x = *&y;
	y = t;

	fmt.Println(x,y);
}

func flagParseCliArgs(){
	var mode = flag.String("mode","","process mode");
	flag.Parse();
	fmt.Println("mode:",mode,"*mode:",*mode);
}

func newPointerType(){
	str := new(string);
	*str = "Learn go!";
	fmt.Println("str:",str,"*str:",*str,"&str:",&str);
}