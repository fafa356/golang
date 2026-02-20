package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := []int{12, 23, 41, 67, 83, 22, 233, 100, 39, 21}
	res := []int{}

	for i := 0; i < 10; i++ {
		sum := arr[i] + slice[i]
		res = append(res, sum)
	}
	fmt.Println(res)

}
