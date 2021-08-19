package int

import (
	"fmt"
	"unsafe"
)

/*

整型(整数类型):  有符号和无符号的整数类型


有符号数和无符号数

有符号数: 表示特定类型规定范围内的整数(包括负数)
无符号数: 表示特定类型规定范围内的非负数

注意: 有符号数能够表示负数的代价是将能够表示范围的约一般用来表示负数

例如:

16位无符号整数可表示 0~65535 之间的数值
16位有符号整数可表示 -32768~32767之间的数值


有符号整型: int int8 int16 int32 int64

无符号整型: uint uint8 uint16 uint32 uint64

rune: 表示 Unicode 字符, 和 int32 等价

byte: 强调数值是原始数据,而不是小的整数, 和 uint8 等价

uintptr: 无符号整数类型, 没有指定具体的 bit 大小, 但足以容纳指针

*/

func main(){
	printInt();
	printUint();
}

func printInt(){
	var ii int = -1;
	var i int = 1;
	var i2 int8 = 1;
	var i3 int16 = 1;
	var i4 int32 = 1;
	var i5 int64 = 1;
	fmt.Println("print int start")
	fmt.Println(unsafe.Sizeof(ii)) // 8
	fmt.Println(unsafe.Sizeof(i)) // 8
	fmt.Println(unsafe.Sizeof(i2)) // 1
	fmt.Println(unsafe.Sizeof(i3)) // 2
	fmt.Println(unsafe.Sizeof(i4)) // 4
	fmt.Println(unsafe.Sizeof(i5)) // 8
	fmt.Println("print int end")
}

func printUint(){
	var i uint = 1;
	var i2 uint8 = 1;
	var i3 uint16 = 1;
	var i4 uint32 = 1;
	var i5 uint64 = 1;
	fmt.Println("print uint start")
	fmt.Println(unsafe.Sizeof(i)) // 8
	fmt.Println(unsafe.Sizeof(i2)) // 1
	fmt.Println(unsafe.Sizeof(i3)) // 2
	fmt.Println(unsafe.Sizeof(i4)) // 4
	fmt.Println(unsafe.Sizeof(i5)) // 8
	fmt.Println("print uint end")
}


/*

计算机存储单元

计算机内部,信息采用二进制形式存储,运算,处理和传输

信息存储单位有位,字节,字等几种

存储设备存储容量单位有KB,MB,GB,TB等几种


位(bit): 二进制数中的一个位数, 0或1,是计算机中数据的最小单位;

	字节(Byte,B): 计算机中数据的基本单位, 每8位(bit)组成一个字节
	
	ASCIIS码: 1个英文字母(不区分大小写) = 1 字节空间
				1个中文汉字 = 2 字节空间
				1个ASCII码 = 1 字节空间

	UTF-8编码: 1个英文字符 = 1 字节空间
				1个英文标点 = 1 字节空间
				1个中文汉字 = 3 字节空间
				1个中文标点 = 3 字节空间

	Unicode编码: 1个英文字符 = 2 字节空间
				1个英文标点 = 2 字节空间
				1个中文汉字 = 2 字节空间
				1个中文标点 = 2 字节空间

	字(Word):  2 字节 = 一个字 (汉字存储单位都是一个字)



扩展的存储单位

1024B = 1K(千)B
1024KB = 1M(兆)B
1024MB = 1G(吉)B
1024GB = 1T(太)B


表示范围和字节大小

-2^(n-1) 到 2^(n-1)-1

int8, 占 1 字节, 表示 -128 ~ 127
int16, 占 2 字节, 表示 -32768 ~ 32767
int32, 占 4 字节, 表示 -2147483648 ~ 2147483647  
int64, 占 8 字节, 表示 -9223372036854776000 ~ 9223372036854775999

uint8, 占 1 字节, 表示 0 ~ 255
uint16, 占 2 字节, 表示 0 ~ 65535
uint32, 占 4 字节, 表示 0 ~ 4294967295
uint64, 占 8 字节, 表示 0 ~ 18446744073709551999

*/