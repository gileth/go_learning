package main

import "fmt"

func main() {
	sliceDefine()
	sliceExpress()
	makeSlice()
	appendSlice()
	zuoye()
}

// 切片申明
func sliceDefine() {
	var a []string
	var b = []int{}
	var c = []bool{false, true}
	fmt.Printf("切片 a 的值：%v, b 的值：%v, c 的值：%v", a, b, c)
	fmt.Println()
}

func sliceExpress() {
	a := []int{1, 2, 3, 5, 6}
	s := a[1:3]
	fmt.Printf("从数组a :%v 索引1至3取值赋给切片，s:%v len(s):%v cap(s):%v\n", a, s, len(s), cap(s))
	s = nil
	s = a[1:]
	fmt.Printf("从数组a :%v 索引为1开始，到最后位置索引取值赋值给切片：%v\n", a, s)
	s = nil
	s = a[:]
	fmt.Printf("将数组a : %v赋值给切片，切片：%v\n", a, s)

}

func makeSlice() {
	a := make([]int, 2, 10)
	//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断
	fmt.Printf("切片a： %v, 长度：%v, 容量：%v\n", a, len(a), cap(a))
	//切片赋值是引用方式，切片被赋值给别的切片，修改别的切片会同时修改原始切片
	a2 := a
	a2[0] = 1
	fmt.Printf("切片a: %v, 切片a2: %v", a, a2)

	//遍历切片
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	for index, value := range a {
		fmt.Printf("for range方式遍历：index : %v, value: %v\n", index, value)
	}
}

func appendSlice() {
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v, len: %v, cap: %v, ptr: %p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	//一次追加多个元素
	numSlice = append(numSlice, 12, 222, 56)
	fmt.Printf("追加后的切片：%v\n", numSlice)
}

func zuoye() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Printf("切片a : %v, len: %v, cap: %v\n", a, len(a), cap(a))
}
