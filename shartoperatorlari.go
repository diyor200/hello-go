package main

import (
	"fmt"
	"time"
)

func main_shart() {
	// testIfElse()
	// testSwitch1()
	// testSwitch2()
	testSwitch3()
}

func testIfElse() {
	points := 20
	if points < 50 {
		fmt.Println("Silver")
	} else if points < 100 {
		fmt.Println("Gold")
	} else {
		fmt.Println("Bronze")
	}
}

func testSwitch1() {
	weekday := time.Now().Weekday()
	fmt.Println(weekday)
	switch weekday {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Xato")
	}
}

func testSwitch2() {
	greeting := 2
	switch {
	case greeting == 1:
		fmt.Println("Assalomu alaykum!")
	case greeting == 2:
		fmt.Println("Privet")
	case greeting == 3:
		fmt.Println("Hi")
	}
}
func testSwitch3() {
	var userChoice string = "to'rt"
	switch userChoice {
	case "bir":
		fmt.Println("C#")
	case "ikki", "uch":
		fmt.Println("Go")
	case "to'rt", "besh", "olti":
		fmt.Println("Python")
	}
}
