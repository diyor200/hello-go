package main

import "fmt"

// Sintaksis - map[KeyType]ValueType

func main_map() {
	statuses := make(map[string]int)

	// map'ga qiymatlar qo'shish
	statuses["pending"] = 0
	statuses["inited"] = 1
	statuses["running"] = 2
	statuses["timedout"] = 3
	statuses["successful"] = 4
	statuses["failed"] = 5
	// mapdan ma'lumot o'qish
	var successfulStatus = statuses["successful"]
	fmt.Println(successfulStatus)

	// mapdan ma'lumot o'chirish
	delete(statuses, "failed")
	fmt.Println(statuses)
	testFor()
}
