package main

import (
	"fmt"
	"time"
)

type calculation struct {
	number, result int
}

func workerpool_func() {
	const jobCount = 40
	jobs := make(chan calculation, jobCount)
	results := make(chan calculation, jobCount)
	start := time.Now()

	for w := 1; w <= 4; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= jobCount; j++ {
		calc := calculation{j, 0}
		jobs <- calc
	}
	close(jobs)

	for a := 1; a <= jobCount; a++ {
		result := <-results
		fmt.Println("Fib(", result.number, ")=", result.result)
	}
	close(results)

	duration := time.Since(start)
	fmt.Println("Ish", duration, "vaqtda tugatildi!")
}

func worker(id int, jobs <-chan calculation, results chan<- calculation) {
	for j := range jobs {
		fmt.Println(id, "chi ishchi", j.number, "chi ishni boshladi")
		j.result = fibonacci(j.number)
		results <- j
		fmt.Println(id, "chi ishchi", j.number, "ni tugatdi.")
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
