package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	point := Point{1, 2}
	a = point

	var b Point
	b = a.(Point)
	//b = a
	fmt.Println(b)

	var x interface{}
	b2 := 1.1
	x = b2
	y := x
	fmt.Printf("y的类型是 %T 值是 %v", y, y)
}
