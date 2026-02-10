package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var y float64
	x := rand.Int31n(100)

	switch {
	case x >= 10 && x <= 50:
		y = float64(x-4) / float64(x) //$y = (x - 4)/x
	default:
		y = float64(x*x) / float64(4+x*x) //$y = x^2 / (4 + x^2)
	}

	fmt.Printf("Значення x = %d\n", x)
	fmt.Printf("Значення у = %.2f\n", y)
}
