package main

func tailCall(n int) int {
	return fibTailCall(n, 0, 1)
}

func fibTailCall(n int, a int, b int) int {
	if n == 0 {
		return a
	}
	return fibTailCall(n-1, b, a+b)
}
