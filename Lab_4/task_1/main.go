package main

import "fmt"

func main() {
	chGenerate := generate()
	chFilterEven := filterEven(chGenerate)
	chSquare := square(chFilterEven)
	sum(chSquare)
}

func generate() chan int {
	nums := make(chan int, 100)

	go func() {
		for i := 1; i <= 100; i++ {
			nums <- i
		}
		close(nums)
	}()

	return nums
}

func filterEven(nums chan int) chan int {
	evenNums := make(chan int)

	go func() {
		for e := range nums {
			if e%2 == 0 {
				evenNums <- e
			}
		}
		close(evenNums)
	}()
	return evenNums
}

func square(evenNums chan int) chan int {
	squareNums := make(chan int)

	go func() {
		for sq := range evenNums {
			squareNums <- sq * sq
		}
		close(squareNums)
	}()
	return squareNums
}

func sum(squares chan int) {
	sumNums := make(chan int, 1)
	total := 0

	go func() {
		for s := range squares {
			total += s
		}
		sumNums <- total
		close(sumNums)
	}()
	fmt.Println(<-sumNums)
}
