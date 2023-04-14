package main

func forLoop(n int) int {
	if n == 0 {
		return 0
	}
	a := 0
	b := 1
	for n > 1 {
		a, b = b, a+b
		n--
	}
	return b
}
