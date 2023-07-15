package main

import "fmt"

func main() {
	pointer1()
	fmt.Println("====================")
	pointerType()
	fmt.Println("====================")
	newFunc()

}

func newFunc() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*a)
	fmt.Println(*b)
}

func pointerType() {
	a := 10
	b := &a
	fmt.Printf("type of b:%T\n", b)
	c := *b
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v", c)
}

func pointer1() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	fmt.Println(&b)
}
