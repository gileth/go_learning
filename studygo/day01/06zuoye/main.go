package main

import "fmt"

//1. 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
//2. 编写代码统计出字符串"hello沙河小王子"中汉字的数量。

func zuoye01() {
	varint := 11
	varfloat := 1.222
	varbool := true
	varstring := "sdfsdkdk是的"

	fmt.Printf("varint 类型：%T, 值：%d \n", varint, varint)
	fmt.Printf("varfloat 类型：%T, 值：%.3f \n", varfloat, varfloat)
	fmt.Printf("varbool 类型：%T，值：%t \n", varbool, varbool)
	fmt.Printf("varstring 类型：%T，值：%v \n", varstring, varstring)

}

func main() {
	zuoye01()
}
