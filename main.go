package main

import (
	"fmt"
	"time"
)

func startPolling() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		fmt.Println("poll happened at", time.Now().Format("15:04:05"))
	}
}	

func main() {
	go startPolling()

	fmt.Println("main is doing other things...")
	time.Sleep(5 * time.Second)
	fmt.Println("main is done, program exits (goroutine dies with it)")
}