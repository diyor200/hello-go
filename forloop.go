package main

import (
	"fmt"
)

func testFor() {
	// 1 - ishlatish usuli
	// for i := 1; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// 2 - ishlatish usuli (infinite ishledi)
	// for {
	// 	fmt.Println("Alhamdulillah ", strconv.Itoa(time.Now().Second()))
	// 	time.Sleep(1 * time.Second)
	// 	break // chiqish uchun ishlatiladi
	// }

	// 3-ishlatish usuli (while ga o'xshash)
	// k := 1
	// for k < 10 {
	// 	fmt.Println(strconv.Itoa(k))
	// 	k++
	// }

	// 4-ishlatish usuli
	// myarr := [3]string{"oy", "quyosh", "yer"}
	mymap := map[int]string{
		1: "c#",
		2: "python",
		3: "go",
	}
	for i, val := range mymap {
		fmt.Println(i, val)
	}
}
