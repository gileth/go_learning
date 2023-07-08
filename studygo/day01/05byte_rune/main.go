package main

import (
	"fmt"
	"math"
)

//https://www.liwenzhou.com/posts/Go/datatype/

// 遍历字符串
func traversalString() {
	s := "hello 沙河"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()

	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()
}

// 修改字符串
func changeString() {
	s1 := "big"
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

// 类型转换
func sqrtDemo() {
	var a, b = 3, 4
	c := int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	traversalString()
	changeString()
	sqrtDemo()
}
