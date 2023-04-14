package main

import "fmt"

func queue_func() {
	testQueue()
}

func testQueue() {
	queue := NewQueue()
	queue.Enqueue(100)
	fmt.Println("queue", queue)
	queue.Enqueue(25).Enqueue(35)
	fmt.Println("After enqueued:", queue)
	fmt.Println("Is empty?", queue.IsEmpty())

	var result, _ = queue.Peek()
	fmt.Println("the next item in the queue:", result)

	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeue()
	fmt.Println(result)
	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeue()
	fmt.Println(result)

	fmt.Println("queue", queue)

}

// Queue turini e'lon qilamiz
type Queue struct {
	data []int
}

// NewQueue : yangi Queue obyektini qaytarib beradi
func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

// IsEmpty : Queue bo'shligini tekshiradi
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek funksiyasi yordamida queue dagi keyingi element aniqlab olamiz
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue bo'sh")
	}
	return q.data[0], nil
}

// Enqueue : queue ga yangi element qo'shadi va queue ga ko'rsatkich qaytarib beradi
func (q *Queue) Enqueue(n int) *Queue {
	q.data = append(q.data, n)
	return q
}

// Dequeue : queue dan 1-elementni oladi va queuedan o'chiradi
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue bo'sh!")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}
