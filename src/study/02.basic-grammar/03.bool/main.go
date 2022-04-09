package main

import (
	"fmt"
)

/*

布尔类型的值只有两种: true, false

1. if 和 for语句条件部分都是布尔类型
2. ==, <, >, 等比较操作产生布尔值
3. 一元操作符!反转bool类型的值
4. &&（AND）和 ||（OR）操作符返回布尔类型的值
5. 布尔型无法参与数值运算，也无法与其他类型进行转换。


*/

func main() {
	var bol bool = false
	fmt.Println("bool:", bol)
}
