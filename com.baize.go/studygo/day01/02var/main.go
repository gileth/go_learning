package main

import "fmt"

var intvar int
var strvar string
var boolvar bool

var (
	a int
	b string
	c bool
	d float32
)

const (
	c1 = iota
	c2
	c3
	_
	c5
	c6 = 200
	c7 = iota
	c8
)
const c9 = iota

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)


func foo() (int, string) {
	return 1, "foobar"
}

func main() {
	intvar := 100
	strvar, boolvar = "sdff", true
	a, b, c, d = 1, "a", true, 1.222
	fmt.Println(intvar, strvar, boolvar, a, b, c, d)

	fooint, _ := foo()
	_, foostr := foo()

	fmt.Println("fooint, ", fooint)
	fmt.Println("foostr, ", foostr)
	fmt.Println(c1, c2, c3, c5, c6, c7, c8, c9)

	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Printf("1kb=%d Byte, 1mb=%d Byte, 1gb=%d Byte, 1tb=%d Byte, 1pb=%d Byte", KB, MB, GB, TB, PB)
}
