package main

import (
	"fmt"
)

func cal (x, y int64) (plus int64) {
	plus = x + y
	return
}

func main2() {
	fmt.Println(cal(20, 10))
}
