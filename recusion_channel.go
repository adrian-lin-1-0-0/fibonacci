package main

func recursionChannel(n int) int {
	if n == 0 {
		return 0
	}
	result := make(chan int)
	fibRecursionChannel(n, 0, 1, result)
	return <-result
}

func fibRecursionChannel(n int, a int, b int, result chan int) {
	if n == 0 {
		result <- a
		return
	}

	go fibRecursionChannel(n-1, b, a+b, result)
}
