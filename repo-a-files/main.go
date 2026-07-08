package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("hello", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main () {
	go sayHello() // run in parallel

	fmt.Println("main continues immediately")
	time.Sleep(2* time.Second) // waiting to see the goroutines result
	fmt.Println("main done")
}