package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main() {
	var c chan string = make(chan string)
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)

	select {
	case msg1 := <-c1:
		fmt.Println("Message 1", msg1)
	case msg2 := <-c2:
		fmt.Println("Message 2", msg2)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
