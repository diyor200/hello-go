package main

import (
	"fmt"
	"strconv"
)

func pointer_func() {
	var x int
	x = 10
	fmt.Println(x)
	fmt.Println(&x) // print address of x

	// pointerni e'lon qilish
	var num *int
	val := new(int)

	num = new(int)
	*num = x // qiymat berish
	val = &x // address berish

	print("===pointer num===")
	print(strconv.Itoa(*num))
	fmt.Println(num)
	print("===pointer value===")
	fmt.Println(*val)
	fmt.Println(val)
	x = 1
	y := 2
	fmt.Println(x, y)
	change_val(&x, &y)
	fmt.Println(x, y)
}

func change_val(num1, num2 *int) {
	var o int
	o = *num1
	*num1 = *num2
	*num2 = o
}

func print(text string) {
	fmt.Println(text)
}
