package main

import "fmt"

// 有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
func main() {
	varint := []int{11, 11, 22, 22, 23, 33, 33, 44, 44}
	result := 0
	for _, num := range varint {
		result ^= num
	}
	fmt.Printf("数组 %v 中只出现一次的数字是：%v", varint, result)
}
