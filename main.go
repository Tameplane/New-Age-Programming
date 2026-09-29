package main

import "fmt"

func square(n int, ch chan int) {
	fmt.Println("Square goroutine started")
	ch <- n * n
	fmt.Println("Square goroutine completed")
}

func cube(n int, ch chan int) {
	fmt.Println("Cube goroutine started")
	ch <- n * n * n
	fmt.Println("Cube goroutine completed")
}

func fibonacci(n int, ch chan int) {
	fmt.Println("Fibonacci goroutine started")
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	ch <- a
	fmt.Println("Fibonacci goroutine completed")
}

func main() {
	result := make(chan int)
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	go square(n, result)
	go cube(n, result)
	go fibonacci(n, result)

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
