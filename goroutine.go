package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutine lar haqida
// WaitGroup lar haqida
func goroutine_func() {
	var wg sync.WaitGroup
	wg.Add(2)
	// 2 nechta goroutine funksiya borligini bildiradi
	go func() {
		display("Salom berdik")
		wg.Done()
	}() // go - goroutine(asinxron) tarzda ishledi
	go func() {
		display("Alik oldik")
		wg.Done()
	}()

	wg.Wait() // waitgropuni countri 0 ga boguncha bloklab turadi
	// 0 ga teng bo'gandan keyin main funksiyasi bajariladi
	fmt.Println("main funksiyasi ishladi")

	// fmt.Scanln() // foydalanuvchidan enterni bosishini kutadi
}

// waitgroup ga qoshilgan funksiya tugallangan keyin
// counter bittaga kamaytirish uchun wg.Done() ni funksiyaga kiritish
// kerak
// func display(input string, wg *sync.WaitGroup) {
func display(input string) {
	for i := 1; i < 5; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}
