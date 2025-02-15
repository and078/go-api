package fibonacci

func Fibonacci (n uint64, ch chan uint64) {
	res := fibonacci(n)
	ch <- res
}

func fibonacci(n uint64) uint64 {
	if n <= 1 {
		return n
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

