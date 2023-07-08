package main

import (
	"fmt"
	"math"
)

func main() {

	//十进制
	var a int = 10
	fmt.Printf("十进制：%d \n", a)
	//二进制
	fmt.Printf("二进制：%b \n", a)

	//八进制
	var b int = 0777
	fmt.Printf("八进制：%o \n", b)

	//十六进制
	var c int = 0xff
	fmt.Printf("十六进制：%x \n", c)
	fmt.Printf("大写十六进制:%X \n", c)

	fmt.Printf("%f \n", math.Pi)
	fmt.Printf("%.2f \n", math.Pi)

	var c1 complex128
	var c2 complex64
	c1 = 1 + 2i
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
}
