package main

func recursionTree(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return recursionTree(n-1) + recursionTree(n-2)
}
