package main

import (
	"fmt"
	"math"
)

/*

Go语言支持两种精度浮点数: float32 和 float64

float32: 1.4e-45 ===> 3.4e38  提供大约6个十进制数的精度
float64:  4.9e-324 ===> 1.8e308  提供大约15个十进制数的精度

通常,应优先使用float64类型, float32类型累计计算误差容易扩散



*/

func main(){
	float32Error();
	floatDeclara();
	printFixedNumber();
}


// float32 累计计算容易产生误差
func float32Error (){
	println("\nfloat32Error-----------------------");
	var f float32 = 17776211;
	fmt.Println(f==f+1);
}

// 浮点数声明可以只写整数部分或者小数部分
// 很小或很大的数建议使用科学计数法
func floatDeclara(){
	println("\nfloatDeclara-----------------------");
	var f1 = .7777;
	var f2 = 1.;
	fmt.Println(f1,f2);

	const f3 = 8.872367476457e23;
	const f4 = 8.872367476457e-23;
	fmt.Println(f3,f4);
}

// 打印浮点数,利用 %f 控制保留几位小数
func printFixedNumber(){
	println("\nprintFixedNumber-----------------------");
	const f = math.Pi;
	fmt.Println(f);
	fmt.Printf("%f\n",f);
	fmt.Printf("%.3f\n",f);
	fmt.Printf("%.2f\n",f);
	fmt.Printf("%.1f\n",f);
}
