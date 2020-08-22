package main

import (
	"fmt"
	"time"
)

func main() {
	chanConcurrent()
}
func chanConcurrent() {
	ch := make(chan int, 1)
	go produceChan(ch, "a")
	go produceChan(ch, "b")
	consumerChan(ch)
}
func consumerChan(ch chan int) {
	i := 0
	for {
		select {
		case <-time.After(time.Millisecond * 100):
			ch <- i
			i++
			ch <- i
			i++
		}
	}
}
func produceChan(ch chan int, name string) {
	for {
		select {
		case c := <-ch:
			fmt.Printf("%v get %v\n", name, c)
		}
	}
}
