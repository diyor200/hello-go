package main

import (
	"fmt"
	"log"
	"os"
)

func logging_func() {
	simpleLogging()
	formattingLogging()
}

func simpleLogging() {
	fmt.Println("-----simple logging-----")
	log.Println("hello world")
	log.Println("This is a simple error")
}

func formattingLogging() {
	fmt.Println("---formatting logging---")
	fmt.Println("---writing file ---")
	file, err := os.OpenFile("myapp.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file")
		return
	}

	var warning *log.Logger

	warning = log.New(
		// os.Stdout, // write to console
		file, // write to log file
		"APP: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	warning.Println("This is waning messge 1")
	warning.Println("This is warning message 2")
	warning.Println("This is warning message 2")
	fmt.Println("Done")
}
