package main

import "fmt"

type rect struct {
	width, height float64
}

func (r *rect) area() float64 {
	r.width = 8
	return r.width * r.height
}

func (r rect) perim() float64 {
	//如果方法的接收者是值类型，无论调用者是对象还是对象指针，修改的都是对象的副本，不影响调用者；如果方法的接收者是指针类型，则调用者修改的是指针指向的对象本身。
	r.width = 12
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("r.width:", r.width)
	fmt.Println("perim:", r.perim())
	fmt.Println("r.width:", r.width)

	//通过指针调用
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim", rp.perim())
}

//以下是运行结果
//area: 40
//r.width: 8
//perim: 34
//r.width: 8
//area: 40
//perim 34
