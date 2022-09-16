package main

import (
	"fmt"
	"math"
)

// создаем структуру
type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Circle struct {
	x1, y1, r1 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func main() {
	r := Rectangle{0, 0, 10, 10}
	c1 := Circle{0, 0, 5}
	fmt.Println(r.area())
	fmt.Println(c1)
}
