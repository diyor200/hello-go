package main

import "fmt"

func main_hello() {
	a := 14
	b := 30

	// & (bitwise AND)
	result1 := a & b
	fmt.Printf("a & b = %d", result1)

	// | (bitwise OR)
	result2 := a | b
	fmt.Printf("\na | b = %d", result2)

	// ^ (bitwise XOR)
	result3 := a ^ b
	fmt.Printf("\na ^ b = %d", result3)

	// << (left shift)
	result4 := a << b
	fmt.Printf("\na << b = %d", result4)

	// misc operatorlari
	x := 4
	// xotiradagi manzilini ishlatish va ko'rsatkich(*) operatorlari
	y := &x
	fmt.Println(*y)
	*y = 7
	fmt.Println(x)
}
