package main

import (
	"fmt"
	"task_2/calc"
)

func main() {
	myCalc := calc.Calc{}

	nums := []float64{1, 2, 3, 2.1, -0.2}

	sum := myCalc.Sum(nums...)
	max := myCalc.Max(nums...)
	min := myCalc.Min(nums...)
	div, err := myCalc.Divide(max, min)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("sum = %.2f\nmax = %.2f\nmin = %.2f\ndiv = %.2f\n", sum, max, min, div)

}
