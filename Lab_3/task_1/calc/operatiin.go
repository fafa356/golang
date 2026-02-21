package calc

import (
	"errors"
	"fmt"
)

func Sum(nums ...float64) (sum float64) {
	for _, num := range nums {
		sum += num
	}
	return
}

func Max(nums ...float64) (max float64) {
	if len(nums) == 0 {

		return 0
	}
	max = nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return
}

func Min(nums ...float64) (min float64) {
	if len(nums) == 0 {
		return 0
	}
	min = nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("ділення на нуль")
	}

	return a / b, nil
}
func init() {
	fmt.Println("Пакет calc успішно ініціалізовано. Готові до обчислень!")
}
