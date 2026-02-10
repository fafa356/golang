package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var y float64

	x := rand.Int31n(100)

	if x >= 1 && x <= 50 {
		y = float64(x-4) / float64(x)
	} else {
		y = float64(x*x) / float64(4+x*x)
	}

	fmt.Printf("Значення x = %d\n", x)
	fmt.Printf("Значення y = %.2f\n", y)

}
