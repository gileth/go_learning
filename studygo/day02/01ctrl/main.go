package main

import (
	"fmt"
)

func main() {
	ifDemo()
	forDemo1()
	forDemo2()
	gotoDemo()
}

func ifDemo() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	fmt.Println("==========================")
}

func forDemo1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func gotoDemo() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakFlag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
breakFlag:
	fmt.Println("退出循环")
}
