package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format(time.DateTime))
	fmt.Println(time.Now().Format("2017-09-07 18:05:32"))

}
