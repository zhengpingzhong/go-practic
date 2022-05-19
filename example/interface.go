package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area2() float64
	perim2() float64
}

type rect2 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect2) area2() float64 {
	return r.width * r.height
}

func (r rect2) perim2() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area2() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim2() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area2())
	fmt.Println(g.perim2())
}

func main() {
	// package main 之前有个method测试里面包含rect 及其area 和 perim方法
	r := rect2{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
