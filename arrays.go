package main

import (
	"fmt"
	"sort"
)

func main_arrays() {
	testArrays1()
	testArrays2()
	testArrays3()
	testArrays4()
	testArrays5()
	testSlice1()
	testSlice2()

}

func testSlice2() {
	myarr := []int{15, 14, 19, 5, 3, 16}
	sort.Ints(myarr)
	fmt.Println(myarr)

}

func testSlice1() {
	my_slice := []int{1, 2, 3}
	my_slice = append(my_slice, 4)
	fmt.Println("Slice ning uzunligi: ", len(my_slice))
}

func testArrays1() {
	var myarr [3]string
	// massivga qiymat berish
	myarr[0] = "Abdulaxatov"
	myarr[1] = "Diyorbek"
	myarr[2] = "Ulug'bek o'g'li"
	fmt.Println(myarr)
}

func testArrays2() {
	// massivga qiymat berishning qisqa usuli
	myarr := [3]int{1, 2, 3}
	fmt.Println(myarr)
}

func testArrays3() {
	myarr := [...]int{2, 3, 4, 5}
	myarr1 := [4]int{2, 3, 4, 5}
	fmt.Println(myarr == myarr1)
}

func testArrays4() {
	myarr := [3][3]string{{"C#", "GO", "Python"},
		{"Java", "Javascript", "Typescript"},
		{"Postgresql", "Mysql", "Oracle"}}
	fmt.Println(myarr[0][2])
}

func testArrays5() {
	// qatordan reference li nusxa olish
	myarr := [3]int{1, 2, 3}
	myarr1 := &myarr
	fmt.Println(myarr)
	fmt.Println(*myarr1)
	myarr[1] = 0
	fmt.Println(myarr)
	fmt.Println(*myarr1)
}
