package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d : fizzbuzz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d : buzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d : fizz\n", i)
		}
	}

}
