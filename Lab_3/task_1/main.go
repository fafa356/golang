package main

import (
	"fmt"
	"task_1/calc"
)

func main() {
	nums := []float64{1, 4, 2, 7, 34, 9, -2, -1, -4}
	sum := calc.Sum(nums...)
	max := calc.Max(nums...)
	min := calc.Min(nums...)
	div, err := calc.Divide(4, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("nums = %v\n sum = %.2f\n max = %.2f\n min = %.2f\n div = %.2f\n ", nums, sum, max, min, div)

}
