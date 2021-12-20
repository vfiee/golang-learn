package main

import (
	"fmt"
)

/*

复数: 由两个浮点数表示,其中一个表示实部(real), 一个表示虚部(imag)

组成: RE+IMi RE表示实数部分 IM表示虚数部分 i表示虚数单位

类型: complex128(64位实数和虚数|复数的默认类型) complex64(32位实数和虚数)

声明:

var name complex128 = complex(x,y);

name := complex(x,y);

x: 表示复数实数部分

y: 表示实数的虚数部分


获取实数和虚数:

var plural complex64 = complex(10,2);

var _real = real(plural);

var _imag = imag(plural);

*/

func main(){
	var plural complex64 = complex(10,2);

	var _real = real(plural);

	var _imag = imag(plural);

	fmt.Println(_real);
	fmt.Println(_imag);
}