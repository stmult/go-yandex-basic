package main

import "fmt"

func main() {
	xb := []float64{30, 93, 77, 82, 83}
	fmt.Println(average(xb))
	fmt.Println(factorial(3))
	fmt.Println(f1())
	fmt.Println(add(1, 2, 3, 4, 5))

	// пример замыкания функции внутри функции
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1100, 1))

	// после паники идет остановка
	panic1(xb)

}

// просто функция, которая считает среднее значение
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// рекурсия
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

// пример вызова паники
func panic1(xs []float64) float64 {
	panic("Not Implemented")
}

// показан стек функций
func f1() int {
	return f2()
}
func f2() int {
	return 1
}

// можно указать через ... чтобы передавалось любое количество параметров, даже 0
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
