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
	half(5)
	max(1, 2, 3)
	max2([]int{100, 200, 300, 500})

	//defer перемещает вызов second в конец функции, но перед паникой
	defer second()
	first()

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // 1
	fmt.Println(nextOdd()) // 3
	fmt.Println(nextOdd()) // 5

	fmt.Println(fibonachi(30))
	// так нужно вызывать панику, если хотим получить сообщение паники. То есть обязательно использовать defer, иначе паника всё остановит
	/* defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC") */

	// после паники идет остановка
	panic1(xb)

}
func max(args2 ...int) {
	start := args2[0]
	for _, i3 := range args2 {
		if i3 > start {
			start = i3
		}
	}
	fmt.Println(start)
}

func max2(xn []int) {
	start := xn[0]
	for _, i3 := range xn {
		if i3 > start {
			start = i3
		}
	}
	fmt.Println(start)
}

func half(x2 int) {
	x3 := x2 / 2
	res := "false"
	if x2%2 == 0 {
		res = "true"
	}
	fmt.Printf("%d, %s\n", x3, res)
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

func fibonachi(x uint) uint {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fibonachi(x-1) + fibonachi(x-2)
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

// генерирует четное число
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
