package main

/*
Go字符类型( byte 和 rune )

byte: byte类型,等价于uint8类型,表示一个ASCII字符
rune: rune类型,等价于int32类型,表示一个UTF-8字符

Unicode 包中内置了测试字符的函数，返回值一个布尔值
判断是否为字母：unicode.IsLetter(ch)
判断是否为数字：unicode.IsDigit(ch)
判断是否为空白符号：unicode.IsSpace(ch)

*/