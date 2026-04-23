package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var count atomic.Int64
	oddChan := make(chan int)
	evenChan := make(chan int)
	go Generator(oddChan, evenChan)

	for {
		if oddChan == nil && evenChan == nil {
			break
		}
		select {
		case msg, ok := <-oddChan:
			if !ok {
				oddChan = nil
				continue
			}
			if msg%33 == 0 {
				count.Add(-1)
			}
		case msg, ok := <-evenChan:
			if !ok {
				evenChan = nil
				continue
			}
			if msg%3 == 0 {
				count.Add(1)
			}
		}
	}
	fmt.Println("count:", count.Load())
}

func Generator(oddChan chan int, evenChan chan int) {
	for i := 1; i <= 1000; i++ {
		if i%2 == 0 {
			evenChan <- i
		} else {
			oddChan <- i
		}
	}
	close(oddChan)
	close(evenChan)
}
