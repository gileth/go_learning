package main

import "fmt"

func main() {
	sm, sb := calcAddSub(3, 2)
	fmt.Printf("add : %d, sub : %d", sm, sb)
	fmt.Println("====================")
	sm2, sb2 := calc(10, 5, calcAddSub)
	fmt.Printf("add2 : %d, sub2 : %d", sm2, sb2)
	fmt.Println("====================")
	anonymousFunc()
	fmt.Println("====================")
	bibaoFunc1()

	fmt.Println("====================")
	deferFunc()
	fmt.Println("====================")
}

func calcAddSub(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func calc(x, y int, op func(int, int) (int, int)) (int, int) {
	return op(x, y)
}

func anonymousFunc() {
	//将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(1, 2) //通过变量调用匿名函数

	//自执行函数，匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(12, 23) //不能写在下一行，必须紧接着匿名函数关闭的}

}

// 闭包
func bibaoFunc1() {
	var f = adder1()
	fmt.Println(f(10))
	fmt.Println(f(20))

}

func adder1() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func deferFunc() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
