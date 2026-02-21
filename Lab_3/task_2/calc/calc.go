package calc

import (
	"errors"
	"fmt"
)

type Calculator interface {
	Sum(nums ...float64) float64
	Max(nums ...float64) float64
	Min(nums ...float64) float64
	Divide(a, b float64) (float64, error)
}

type Calc struct {
}

func (c Calc) Sum(nums ...float64) (sum float64) {
	for _, num := range nums {
		sum += num
	}
	return
}

func (c Calc) Max(nums ...float64) (max float64) {
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

func (c Calc) Min(nums ...float64) (min float64) {
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

func (c Calc) Divide(a, b float64) (float64, error) {
	if b == 0 {

		return 0, errors.New("ділення на нуль")
	}

	return a / b, nil
}
func init() {
	fmt.Println("Пакет calc успішно ініціалізовано. Готові до обчислень!")
}
