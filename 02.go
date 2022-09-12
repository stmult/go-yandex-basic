package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x[4])

	var x2 [5]float64
	x2[0] = 98
	x2[1] = 93
	x2[2] = 77
	x2[3] = 82
	x2[4] = 83

	var total float64 = 0
	for i := 0; i < len(x2); i++ {
		total += x2[i]
	}
	fmt.Println(total / float64(len(x2)))

	var total2 float64 = 0
	for _, value := range x2 {
		total2 += value
	}
	fmt.Println(total2 / float64(len(x)))

	x3 := [5]float64{98, 93, 77, 82, 83}
	fmt.Println(x3[2])

	x4 := [5]float64{
		1,
		93,
		77,
		82,
		83,
	}
	fmt.Println(x4[4])
}
