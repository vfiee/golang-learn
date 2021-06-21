package main

/*
数据类型转换

Go语言不存在隐式类型转换,所有类型转换都必须显示声明

实例:
current := 10.0;
intCurrent := int(current);


注意:
1.类型转换只能在定义正确的情况下转换成功,如从一个取之范围较小的类型转换为取值范围较大的类型,当从取值范围较大的类型转换
为取值范围较小的类型时,会发生精度丢失(截断)
2.只有相同底层类型的变量之间可以进行相互转换（如将 int16 类型转换成 int32 类型），
不同底层类型的变量相互转换时会引发编译错误（如将 bool 类型转换为 int 类型）


*/

import (
	"fmt"
	"math"
)

func main() {

        // 输出各数值范围
        fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
        fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
        fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
        fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)

        // 初始化一个32位整型值
        var a int32 = 1047483647
        // 输出变量的十六进制形式和十进制值
        fmt.Printf("int32===>:16进制: 0x%x 10进制: %d\n", a, a)

        // 将a变量数值转换为十六进制, 发生数值截断
        b := int16(a)
        // 输出变量的十六进制形式和十进制值
        fmt.Printf("int16===>: 16进制: 0x%x 10进制: %d\n", b, b)

        // 将常量保存为float32类型
        var c float32 = math.Pi
        // 转换为int类型, 浮点发生精度丢失
        fmt.Println(int(c))
}