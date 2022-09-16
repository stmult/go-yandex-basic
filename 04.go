package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
	fmt.Println(*x)
}
func main() {
	x := 5.0
	square(&x)
	x1 := 1
	y1 := 2
	swap(&x1, &y1)

}

func swap(x1, y1 *int) {
	*x1, *y1 = *y1, *x1
	fmt.Println(*x1, *y1)
}
