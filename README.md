# Fibonacci

## Q-Matrix

$\begin{aligned}
Q &= \begin{bmatrix} 1 & 1 \\ 1 & 0 \end{bmatrix} = \begin{bmatrix} F_2 & F_1 \\ F_1 & F_0 \end{bmatrix}
\end{aligned}$

$\begin{aligned}
Q^{n+1} &= QQ^n \\ 
        &= \begin{bmatrix} 1 & 1 \\ 1 & 0 \end{bmatrix} \begin{bmatrix} F_{n+1} & F_n \\ F_n & F_{n-1} \end{bmatrix} \\
        &= \begin{bmatrix} F_{n+1} + F_n & F_n + F_{n-1} \\ F_{n+1} &F_n \end{bmatrix} \\
        &= \begin{bmatrix} F_{n+2} & F_{n+1} \\ F_{n+1} & F_{n} \end{bmatrix}
\end{aligned}$

## Fast doubling

$\begin{aligned}
\begin{bmatrix} 1 & 1 \\ 1 & 0 \end{bmatrix}^n = \begin{bmatrix} F(n+1) & F(n) \\ F(n) & F(n-1) \end{bmatrix}
\end{aligned}$

$\begin{aligned}
\begin{bmatrix} F(2n+1) \\ F(2n) \end{bmatrix}^n &= \begin{bmatrix} F(n+1) & F(n) \\ F(n) & F(n-1) \end{bmatrix}^{2n} \begin{bmatrix} 1 \\ 0 \end{bmatrix} \\
                                                 &= \begin{bmatrix} F(n+1) & F(n) \\ F(n) & F(n-1) \end{bmatrix} \begin{bmatrix} F(n+1) & F(n) \\ F(n) & F(n-1) \end{bmatrix} \begin{bmatrix} 1 \\ 0 \end{bmatrix} \\
                                                 &= \begin{bmatrix} F(n+1)^2 + F(n)^2 \\ F(n)F(n+1) + F(n-1)F(n) \end{bmatrix}
\end{aligned}$


$\begin{aligned}
F(2n+1)  &= F(n+1)^2 + F(n)^2 \\
\end{aligned}$

$\begin{aligned}
F(2n)  &= F(n)F(n+1) + F(n-1)F(n) \\
       &= F(n)[F(n+1)+F(n-1) \\
       &= F(n)[F(n+1)+F(n-1)+F(n)-F(n)] \\
       &= F(n)[F(n+1)+F(n+1)-F(n)] \\
       &= F(n)[2*F(n+1)-F(n)]
\end{aligned}$

Golang code

```go
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
```


```
BenchmarkFastDoubling-8       	1000000000	         0.0000018 ns/op	       0 B/op	       0 allocs/op
BenchmarkRecursionChannel-8   	1000000000	         0.002087 ns/op	               0 B/op	       0 allocs/op
BenchmarkTailCall-8           	1000000000	         0.0003478 ns/op	       0 B/op	       0 allocs/op
BenchmarkForLoop-8            	1000000000	         0.0000029 ns/op	       0 B/op	       0 allocs/op
```
