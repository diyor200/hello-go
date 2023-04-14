package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channellar haqida
// select statement lar haqida
func channel_func() {
	// channel_func1()
	// channel_func2()
	channel_func3()
}

// 1-misol
func channel_func1() {
	// buffered channel
	channel_buffered := make(chan string, 2) // chan kalit sozi orqali yaratiladi
	channel_buffered <- "anor"               // <- belgi orqali channelga ma'lumot uzatiladi
	channel_buffered <- "olma"
	// result := <-channel // channeldagi ma'lumotni o'qish

	// unbufferred channel
	channel := make(chan string)
	go func() {
		channel <- "anor"
		channel <- "olma"
	}()
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

// 2-misol
func channel_func2() {
	channel := make(chan int)
	go getRandonNumber(channel)
	for number := range channel {
		fmt.Println("tasodifiy son:", number)
	}
}

func getRandonNumber(channel chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		number := rand.Intn(1000)
		time.Sleep(1 * time.Second)
		channel <- number
	}
	close(channel)
}

// 3-misol
func channel_func3() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "tez"
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			channel2 <- "sekin"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)
		case message2 := <-channel2:
			fmt.Println(message2)
		default:
			fmt.Println("ma'lumot yo'q!")
		}
	}
}
