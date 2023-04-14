package main

import "fmt"

// struct haqida

type car struct {
	make, model, color string
	year, weight       int
	Engine             engine
}

type engine struct {
	name string
	hp   int
}

func struct_func() {
	// var myCar car
	myCar := car{make: "Gentra", model: "Sedan", color: "White", year: 2023, weight: 1564, Engine: engine{"1.6 motor", 98}}
	fmt.Println(myCar)
	fmt.Println(myCar.Engine.name)
}
