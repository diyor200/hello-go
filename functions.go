package main

import (
	"errors"
	"fmt"
	"math"
)

func functions() {
	result1 := sum(1, 2)
	fmt.Println(result1)

	result, err := sqrt(-9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	a := 5
	b := 6
	fmt.Println(a, b)
	swap(a, b)
	swapByRef(&a, &b)
	fmt.Println(a, b)
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Manfiy sonning ildizi mavjud emas")
	} else {
		return math.Sqrt(x), nil
	}
}

// funksiyaga argumentlarni uzatish by value
func swap(x, y int) int {
	var o int
	o = x
	x = y
	y = o
	return o
}

// funksiyaga argumentlani uzatish by reference
func swapByRef(x, y *int) int {
	var o int
	o = *x
	*x = *y
	*y = o
	return o

}
