package main

func fastDoubling(n int) int {
	if n == 0 {
		return 0
	}

	fn := 1   //F(n)
	fn1 := 1  //F(n+1)
	f2n := 1  //F(2n)
	f2n1 := 0 //F(2n+1)
	i := 1

	for i < n {
		if (i << 1) <= n {
			f2n1, f2n = fn1*fn1+fn*fn, fn*(2*fn1-fn)
			fn, fn1 = f2n, f2n1
			i = i << 1
		} else {
			fn, f2n = f2n, f2n1
			f2n1 = fn + f2n1 // F(2n) + F(2n+1)
			i++
		}
	}
	return f2n
}
