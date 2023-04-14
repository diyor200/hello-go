package fileio

import (
	"fmt"
	"io/ioutil"
)

const filePath = "C:/Users/dabdu/Desktop/hello/fileio/testgo.txt"

func main() {
	fmt.Println("File yaratilmoqda...")
	writeFile("Assalomu alaykum")
	fmt.Println("Filedan o'qilmoqda...")
	readFile()
}

func writeFile(data string) {
	bytes := []byte(data)
	ioutil.WriteFile(filePath, bytes, 0644)
	fmt.Println("Created file")
}

func readFile() {
	data, _ := ioutil.ReadFile(filePath)
	fmt.Println("file content: ")
	fmt.Println(string(data))
}
